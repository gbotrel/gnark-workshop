package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"math/big"
	"os"
	"os/exec"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gbotrel/gnark-workshop/circuit"
)

var fInit = flag.Bool("init", false, "set to true to run circuit Setup and export solidity Verifier")

const (
	r1csPath     = "circuit/mimc.r1cs"
	pkPath       = "circuit/mimc.pk"
	vkPath       = "circuit/mimc.vk"
	solidityPath = "circuit/mimc_verifier.sol"
)

/*
	Need:
	* install abigen
	* if fInit is set, run circuit Setup and export solidity verifier.
*/
func main() {
	flag.Parse()
	if *fInit {
		initCircuit()
		return
	}

	// check that init was performed
	if _, err := os.Stat(r1csPath); os.IsNotExist(err) {
		log.Fatal("please run with -init flag first to serialize circuit, keys and solidity contract")
	}

	// setup geth simulated backend, deploy smart contract
	verifierContract, err := deploySolidity()
	assertNoError(err)

	// read R1CS, proving key and verifying keys
	r1cs := groth16.NewCS(ecc.BN254)
	pk := groth16.NewProvingKey(ecc.BN254)
	vk := groth16.NewVerifyingKey(ecc.BN254)
	deserialize(r1cs, r1csPath)
	deserialize(pk, pkPath)
	deserialize(vk, vkPath)

	// Now we want to create a valid proof
	// 1. We compute our secret, and the hash of our secret
	// 2. Then, we assign these values to our witness (aka circuit input)
	// 3. Then, we ensure the proof verifies in plain Go
	// 4. Finally, we build the solidity input and submit the transaction to the blockchain.

	// pick a secret
	const secret = "secret"
	// hash it with mimc
	var hash []byte
	hFunc := mimc.NewMiMC("seed")
	hFunc.Write([]byte(secret))
	hash = hFunc.Sum(hash)

	// assign values to witness
	var witness circuit.Circuit
	witness.Hash.Assign(hash)
	witness.Secret.Assign([]byte(secret))

	// create the proof
	log.Println("creating proof")
	proof, err := groth16.Prove(r1cs, pk, &witness)
	assertNoError(err)

	// ensure gnark (Go) code verifies it
	err = groth16.Verify(proof, vk, &witness)
	assertNoError(err)

	// solidity contract inputs
	// a, b and c are the 3 ecc points in the proof we feed to the pairing
	// they are stored in the same order in the golang data structure
	// each coordinate is a field element, of size fp.Bytes bytes
	var (
		a     [2]*big.Int
		b     [2][2]*big.Int
		c     [2]*big.Int
		input [1]*big.Int
	)

	// get proof bytes
	var buf bytes.Buffer
	proof.WriteRawTo(&buf)
	proofBytes := buf.Bytes()

	// proof.Ar, proof.Bs, proof.Krs
	const fpSize = fp.Bytes
	a[0] = new(big.Int).SetBytes(proofBytes[fpSize*0 : fpSize*1])
	a[1] = new(big.Int).SetBytes(proofBytes[fpSize*1 : fpSize*2])
	b[0][0] = new(big.Int).SetBytes(proofBytes[fpSize*2 : fpSize*3])
	b[0][1] = new(big.Int).SetBytes(proofBytes[fpSize*3 : fpSize*4])
	b[1][0] = new(big.Int).SetBytes(proofBytes[fpSize*4 : fpSize*5])
	b[1][1] = new(big.Int).SetBytes(proofBytes[fpSize*5 : fpSize*6])
	c[0] = new(big.Int).SetBytes(proofBytes[fpSize*6 : fpSize*7])
	c[1] = new(big.Int).SetBytes(proofBytes[fpSize*7 : fpSize*8])

	// public witness, the hash of the secret is on chain
	input[0] = new(big.Int).SetBytes(hash)

	// call the contract
	res, err := verifierContract.VerifyProof(nil, a, b, c, input)
	assertNoError(err)

	if !res {
		log.Fatal("calling the verifier on chain didn't succeed, but should have")
	}
	log.Println("successfully verified proof on-chain")

	// (wrong) public witness
	input[0] = new(big.Int).SetUint64(42)

	// call the contract should fail
	res, err = verifierContract.VerifyProof(nil, a, b, c, input)
	assertNoError(err)
	if res {
		log.Println("calling the verifier suceeded, but shouldn't have")
	}

}

func deploySolidity() (*circuit.Verifier, error) {
	const gasLimit uint64 = 8000029
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	auth := bind.NewKeyedTransactor(key)
	genesis := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: big.NewInt(10000000000)},
	}
	simulatedBackend := backends.NewSimulatedBackend(genesis, gasLimit)

	// deploy verifier contract
	log.Println("deploying verifier contract on chain")
	_, _, verifierContract, err := circuit.DeployVerifier(auth, simulatedBackend)
	if err != nil {
		return nil, err
	}
	simulatedBackend.Commit()
	return verifierContract, nil
}

func initCircuit() {
	_, err := exec.LookPath("abigen")
	if err != nil {
		log.Fatal("please install abigen", err)
	}

	var circuit circuit.Circuit

	// compile circuit
	log.Println("compiling circuit")
	r1cs, err := frontend.Compile(ecc.BN254, backend.GROTH16, &circuit)
	assertNoError(err)

	// run groth16 trusted setup
	log.Println("running groth16.Setup")
	pk, vk, err := groth16.Setup(r1cs)
	assertNoError(err)

	// serialize R1CS, proving & verifying key
	log.Println("serialize R1CS (circuit)", r1csPath)
	serialize(r1cs, r1csPath)

	log.Println("serialize proving key", pkPath)
	serialize(pk, pkPath)

	log.Println("serialize verifying key", vkPath)
	serialize(vk, vkPath)

	// export verifying key to solidity
	log.Println("export solidity verifier", solidityPath)
	f, err := os.Create(solidityPath)
	assertNoError(err)
	err = vk.ExportSolidity(f)
	assertNoError(err)

	// run abigen to generate go wrapper
	// abigen --sol circuit/mimc_verifier.sol --pkg circuit --out circuit/wrapper.go
	cmd := exec.Command("abigen", "--sol", solidityPath, "--pkg", "circuit", "--out", "circuit/wrapper.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	assertNoError(err)
}

// serialize gnark object to given file
func serialize(gnarkObject io.WriterTo, fileName string) {
	f, err := os.Create(fileName)
	assertNoError(err)

	_, err = gnarkObject.WriteTo(f)
	assertNoError(err)
}

// deserialize gnark object from given file
func deserialize(gnarkObject io.ReaderFrom, fileName string) {
	f, err := os.Open(fileName)
	assertNoError(err)

	_, err = gnarkObject.ReadFrom(f)
	assertNoError(err)
}

func assertNoError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
