// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package circuit

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PairingABI is the input ABI used to generate the binding from.
const PairingABI = "[]"

// PairingBin is the compiled bytecode used for deploying new contracts.
var PairingBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220350862315ce54faadebd18d61a503b2d252898176feeb844d365a2cb4f509ac564736f6c63430008020033"

// DeployPairing deploys a new Ethereum contract, binding an instance of Pairing to it.
func DeployPairing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pairing, error) {
	parsed, err := abi.JSON(strings.NewReader(PairingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PairingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// Pairing is an auto generated Go binding around an Ethereum contract.
type Pairing struct {
	PairingCaller     // Read-only binding to the contract
	PairingTransactor // Write-only binding to the contract
	PairingFilterer   // Log filterer for contract events
}

// PairingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairingSession struct {
	Contract     *Pairing          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairingCallerSession struct {
	Contract *PairingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PairingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairingTransactorSession struct {
	Contract     *PairingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PairingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairingRaw struct {
	Contract *Pairing // Generic contract binding to access the raw methods on
}

// PairingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairingCallerRaw struct {
	Contract *PairingCaller // Generic read-only contract binding to access the raw methods on
}

// PairingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairingTransactorRaw struct {
	Contract *PairingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairing creates a new instance of Pairing, bound to a specific deployed contract.
func NewPairing(address common.Address, backend bind.ContractBackend) (*Pairing, error) {
	contract, err := bindPairing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// NewPairingCaller creates a new read-only instance of Pairing, bound to a specific deployed contract.
func NewPairingCaller(address common.Address, caller bind.ContractCaller) (*PairingCaller, error) {
	contract, err := bindPairing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairingCaller{contract: contract}, nil
}

// NewPairingTransactor creates a new write-only instance of Pairing, bound to a specific deployed contract.
func NewPairingTransactor(address common.Address, transactor bind.ContractTransactor) (*PairingTransactor, error) {
	contract, err := bindPairing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairingTransactor{contract: contract}, nil
}

// NewPairingFilterer creates a new log filterer instance of Pairing, bound to a specific deployed contract.
func NewPairingFilterer(address common.Address, filterer bind.ContractFilterer) (*PairingFilterer, error) {
	contract, err := bindPairing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairingFilterer{contract: contract}, nil
}

// bindPairing binds a generic wrapper to an already deployed contract.
func bindPairing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.PairingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transact(opts, method, params...)
}

// VerifierABI is the input ABI used to generate the binding from.
const VerifierABI = "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[1]\",\"name\":\"input\",\"type\":\"uint256[1]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VerifierFuncSigs maps the 4-byte function signature to its string representation.
var VerifierFuncSigs = map[string]string{
	"43753b4d": "verifyProof(uint256[2],uint256[2][2],uint256[2],uint256[1])",
}

// VerifierBin is the compiled bytecode used for deploying new contracts.
var VerifierBin = "0x608060405234801561001057600080fd5b5061123a806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806343753b4d14610030575b600080fd5b61004361003e36600461102d565b610057565b604051901515815260200160405180910390f35b6000610061610e0d565b6040805180820182528751815260208089015181830152908352815160808101835287515181840190815288518301516060830152815282518084018452888301805151825251830151818401528183015283820152815180830183528651815286820151918101919091529082015260006100db610586565b6040805180820190915260008082526020820152835151919250906000805160206111e5833981519152116101575760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61582d6774652d7072696d652d7100000000000000000060448201526064015b60405180910390fd5b8251602001516000805160206111e5833981519152116101b95760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61592d6774652d7072696d652d71000000000000000000604482015260640161014e565b602083015151516000805160206111e58339815191521161021c5760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258302d6774652d7072696d652d710000000000000000604482015260640161014e565b6020838101510151516000805160206111e5833981519152116102815760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259302d6774652d7072696d652d710000000000000000604482015260640161014e565b6020838101515101516000805160206111e5833981519152116102e65760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258312d6774652d7072696d652d710000000000000000604482015260640161014e565b60208381015181015101516000805160206111e58339815191521161034d5760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259312d6774652d7072696d652d710000000000000000604482015260640161014e565b6040830151516000805160206111e5833981519152116103af5760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63582d6774652d7072696d652d71000000000000000000604482015260640161014e565b6000805160206111e5833981519152836040015160200151106104145760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63592d6774652d7072696d652d71000000000000000000604482015260640161014e565b60005b6001811015610532577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000186826001811061046157634e487b7160e01b600052603260045260246000fd5b6020020151106104b35760405162461bcd60e51b815260206004820152601f60248201527f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c6400604482015260640161014e565b61051e8261051985608001518460016104cc919061112f565b600281106104ea57634e487b7160e01b600052603260045260246000fd5b602002015189856001811061050f57634e487b7160e01b600052603260045260246000fd5b60200201516108a6565b610942565b91508061052a8161117d565b915050610417565b50608082015151610544908290610942565b905061057a61055684600001516109da565b84602001518460000151856020015185876040015189604001518960600151610a73565b98975050505050505050565b61058e610e5e565b6040805180820182527f0e57d5bc3a479e20053746ce604ed6583f5c8f29cf4aac72293ebe68d3b035c081527f2774101acd932692bd0099db7cad10ab4a0631f472f8037c3eb57b68aa4a244a6020808301919091529083528151608080820184527f06981ef79f49bdeecb54b4f35ec4e28e9fdb44f38c36f44fbcc000db16c40d158285019081527f0dc69638a1c156cc5396d6bd7a78b4b66067dea220f6c67d96139a0de91dd029606080850191909152908352845180860186527f241b12ea8569be207e982fec3b4cc20c1759fb63b636f744390249ad0afee9ce81527f23f645c36385e8b0ed4fd0f2e4f2d499d9615b827f3350e371e53de7542a0e32818601528385015285840192909252835180820185527f2fe125514736fd2b74465c31b83c53d323facef612c1fc79353c042a1047e1ba8186019081527f18951a07bb2f1b7b15337c6827b79668cb95d7be599c4d4809679ea81421d571828501528152845180860186527f28f925ffd2f58180d78fa5e5d6a5c651bd37938d3d265afad0b4f8ebd19a91cb81527f236e2b71cb17064f96be8bd28620cdca534dc767381b47bc55731716a278d878818601528185015285850152835180820185527f255d913dca1714c1ba87f6dc520046cbd6368dd21728bdadb39e545fbf33f1ed8186019081527f08bf8582ffdd64a639c7acc318bde357148831dd1a282b491c9a7cf89008098d828501528152845180860186527f27d12e3a709a8d6bc0bb989b97f2018ae838e9af476ae3027137de7592e279e581527f2ab1095fcaa326d9762b6992d824f2f5949abe5812bd78f785b42330957f77e2818601528185015291850191909152825180840184527f0b2a8e9d8b2f4f347a4bb2ce8f9bbf5c0369a963a5ff271ed0ba9209a4f35d4e81527f05ba6021bad68ffc39cd83576f5fa3d0fa8be857452cd5fa485fe5c1343f1a9f8184015290840180519190915282518084019093527e4066b55d7bd623ccba38abe54955d83975ef80571ab1ce376f73eb5497e37683527f260dff5d6863ab95d0b2f8d1517591eb5789a5d7c1a4da510993a6359c69575e8383015251015290565b60408051808201909152600080825260208201526108c2610eaf565b835181526020808501519082015260408101839052600060608360808460076107d05a03fa90508080156108f5576108f7565bfe5b508061093a5760405162461bcd60e51b81526020600482015260126024820152711c185a5c9a5b99cb5b5d5b0b59985a5b195960721b604482015260640161014e565b505092915050565b604080518082019091526000808252602082015261095e610ecd565b8351815260208085015181830152835160408301528301516060808301919091526000908360c08460066107d05a03fa90508080156108f557508061093a5760405162461bcd60e51b81526020600482015260126024820152711c185a5c9a5b99cb5859190b59985a5b195960721b604482015260640161014e565b604080518082019091526000808252602082015281511580156109ff57506020820151155b15610a1e57506040805180820190915260008082526020820152610a6e565b6040518060400160405280836000015181526020016000805160206111e58339815191528460200151610a519190611198565b610a69906000805160206111e5833981519152611166565b905290505b919050565b60408051608080820183528a825260208083018a90528284018890526060808401879052845192830185528b83528282018a9052828501889052820185905283516018808252610320820190955260009491859190839082016103008036833701905050905060005b6004811015610d8b576000610af2826006611147565b9050858260048110610b1457634e487b7160e01b600052603260045260246000fd5b60200201515183610b2683600061112f565b81518110610b4457634e487b7160e01b600052603260045260246000fd5b602002602001018181525050858260048110610b7057634e487b7160e01b600052603260045260246000fd5b60200201516020015183826001610b87919061112f565b81518110610ba557634e487b7160e01b600052603260045260246000fd5b602002602001018181525050848260048110610bd157634e487b7160e01b600052603260045260246000fd5b6020020151515183610be483600261112f565b81518110610c0257634e487b7160e01b600052603260045260246000fd5b602002602001018181525050848260048110610c2e57634e487b7160e01b600052603260045260246000fd5b6020020151516001602002015183610c4783600361112f565b81518110610c6557634e487b7160e01b600052603260045260246000fd5b602002602001018181525050848260048110610c9157634e487b7160e01b600052603260045260246000fd5b602002015160200151600060028110610cba57634e487b7160e01b600052603260045260246000fd5b602002015183610ccb83600461112f565b81518110610ce957634e487b7160e01b600052603260045260246000fd5b602002602001018181525050848260048110610d1557634e487b7160e01b600052603260045260246000fd5b602002015160200151600160028110610d3e57634e487b7160e01b600052603260045260246000fd5b602002015183610d4f83600561112f565b81518110610d6d57634e487b7160e01b600052603260045260246000fd5b60209081029190910101525080610d838161117d565b915050610adc565b50610d94610eeb565b6000602082602086026020860160086107d05a03fa90508080156108f5575080610df85760405162461bcd60e51b81526020600482015260156024820152741c185a5c9a5b99cb5bdc18dbd9194b59985a5b1959605a1b604482015260640161014e565b505115159d9c50505050505050505050505050565b6040805160a081019091526000606082018181526080830191909152815260208101610e37610f09565b8152602001610e59604051806040016040528060008152602001600081525090565b905290565b6040805160e08101909152600060a0820181815260c0830191909152815260208101610e88610f09565b8152602001610e95610f09565b8152602001610ea2610f09565b8152602001610e59610f29565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060400160405280610f1c610f62565b8152602001610e59610f62565b60405180604001604052806002905b6040805180820190915260008082526020820152815260200190600190039081610f385790505090565b60405180604001604052806002906020820280368337509192915050565b600082601f830112610f90578081fd5b6020610f9b816110dd565b8084868487011115610fab578485fd5b845b6001811015610fca57813584529284019290840190600101610fad565b50909695505050505050565b600082601f830112610fe6578081fd5b610ff060406110dd565b808385604086011115611001578384fd5b835b6002811015611022578135845260209384019390910190600101611003565b509095945050505050565b6000806000806101208587031215611043578384fd5b61104d8686610fd6565b9350604086605f87011261105f578384fd5b600261107261106d8261110e565b6110dd565b8083890160c08a018b811115611086578889fd5b885b858110156110ae5761109a8d84610fd6565b855260209094019391860191600101611088565b508298506110bc8c82610fd6565b97505050505050506110d2866101008701610f80565b905092959194509250565b604051601f8201601f1916810167ffffffffffffffff81118282101715611106576111066111ce565b604052919050565b600067ffffffffffffffff821115611128576111286111ce565b5060200290565b60008219821115611142576111426111b8565b500190565b6000816000190483118215151615611161576111616111b8565b500290565b600082821015611178576111786111b8565b500390565b6000600019821415611191576111916111b8565b5060010190565b6000826111b357634e487b7160e01b81526012600452602481fd5b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a264697066735822122012d33aa517c5732e1b6a178ddc4f00f5079793383ea6d3caa31e7fb3a3af8ad964736f6c63430008020033"

// DeployVerifier deploys a new Ethereum contract, binding an instance of Verifier to it.
func DeployVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifier, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// Verifier is an auto generated Go binding around an Ethereum contract.
type Verifier struct {
	VerifierCaller     // Read-only binding to the contract
	VerifierTransactor // Write-only binding to the contract
	VerifierFilterer   // Log filterer for contract events
}

// VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierSession struct {
	Contract     *Verifier         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierCallerSession struct {
	Contract *VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierTransactorSession struct {
	Contract     *VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierRaw struct {
	Contract *Verifier // Generic contract binding to access the raw methods on
}

// VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierCallerRaw struct {
	Contract *VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierTransactorRaw struct {
	Contract *VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifier creates a new instance of Verifier, bound to a specific deployed contract.
func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// NewVerifierCaller creates a new read-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

// NewVerifierTransactor creates a new write-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

// NewVerifierFilterer creates a new log filterer instance of Verifier, bound to a specific deployed contract.
func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

// bindVerifier binds a generic wrapper to an already deployed contract.
func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_Verifier *VerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_Verifier *VerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}
