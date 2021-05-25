# gnark-workshop

1. Install abigen 
```go
    git clone https://github.com/ethereum/go-ethereum.git 
    cd go-ethereum
    make devtools
```
2. Run `go run main.go -init` to serialize the circuit, its keys and the solidity contract
3. Run `go run main.go` to verify the proof on-chain