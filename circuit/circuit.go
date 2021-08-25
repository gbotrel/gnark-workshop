package circuit

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

// Circuit defines a pre-image knowledge proof
// mimc(secret preImage) = public hash
type Circuit struct {
	Secret frontend.Variable
	Hash   frontend.Variable `gnark:",public"` // struct tags default visibility is "secret"
}

// Define declares the circuit's constraints
// assert mimc(secret) == hash
func (circuit *Circuit) Define(curveID ecc.ID, cs *frontend.ConstraintSystem) error {
	const seed = "seed"

	// hash function
	mimc, err := mimc.NewMiMC(seed, curveID, cs)
	if err != nil {
		return err
	}

	// assert mimc(secret) == hash
	mimc.Write(circuit.Secret)
	cs.AssertIsEqual(mimc.Sum(), circuit.Hash)

	return nil
}

//go:generate go run main.go && abigen --sol solidity/mimc_verifier.sol --pkg solidity --out solidity/solidity.go
