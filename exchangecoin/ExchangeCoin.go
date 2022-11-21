// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exchangecoin

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ExchangecoinMetaData contains all meta data concerning the Exchangecoin contract.
var ExchangecoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"}],\"name\":\"eventReceiverMoney\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"}],\"name\":\"eventWithDrawMoney\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"getEthSignedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveMoney\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ethSignedMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdrawMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405262aa36a760025534801561001757600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506112e5806100676000396000f3fe6080604052600436106100865760003560e01c80639025e64c116100595780639025e64c1461011457806397aba7f91461013f578063a5810c6e1461017c578063a7bb5803146101b9578063fa540801146101f857610086565b806302d05d3f1461008b57806312065fe0146100b65780636c9ee8d1146100e15780636d26ec181461010a575b600080fd5b34801561009757600080fd5b506100a0610235565b6040516100ad9190610703565b60405180910390f35b3480156100c257600080fd5b506100cb610259565b6040516100d89190610737565b60405180910390f35b3480156100ed57600080fd5b50610108600480360381019061010391906109b7565b610261565b005b61011261053d565b005b34801561012057600080fd5b5061012961057c565b6040516101369190610737565b60405180910390f35b34801561014b57600080fd5b5061016660048036038101906101619190610aa0565b610582565b6040516101739190610703565b60405180910390f35b34801561018857600080fd5b506101a3600480360381019061019e9190610b28565b6105f1565b6040516101b09190610bba565b60405180910390f35b3480156101c557600080fd5b506101e060048036038101906101db9190610bd5565b61062a565b6040516101ef93929190610c3a565b60405180910390f35b34801561020457600080fd5b5061021f600480360381019061021a9190610c71565b610692565b60405161022c9190610bba565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600047905090565b81600254146102a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029c90610cfb565b60405180910390fd5b600015156001826040516102b99190610d8c565b908152602001604051809103902060009054906101000a900460ff16151514610317576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030e90610def565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610385576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037c90610e5b565b60405180910390fd5b478311156103c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103bf90610ec7565b60405180910390fd5b60006103d6868686866105f1565b905060006103e382610692565b905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166104268285610582565b73ffffffffffffffffffffffffffffffffffffffff161461047c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047390610f33565b60405180910390fd5b8673ffffffffffffffffffffffffffffffffffffffff166108fc869081150290604051600060405180830381858888f193505050501580156104c2573d6000803e3d6000fd5b50600180846040516104d49190610d8c565b908152602001604051809103902060006101000a81548160ff0219169083151502179055507f33ba68999787348f9a24df6205e34d39f59cbed626f43d7393a29b88706ae94987848660405161052c93929190610ffc565b60405180910390a150505050505050565b7f2ae6f58e7586dcc615a487ca18d0ed0e5721d25d6920b0a12c30f101ac68c6df33346002546040516105729392919061103a565b60405180910390a1565b60025481565b6000806000806105918561062a565b925092509250600186828585604051600081526020016040526040516105ba9493929190611071565b6020604051602081039080840390855afa1580156105dc573d6000803e3d6000fd5b50505060206040510351935050505092915050565b60008484848460405160200161060a9493929190611166565b604051602081830303815290604052805190602001209050949350505050565b60008060006041845114610673576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066a906111fc565b60405180910390fd5b6020840151925060408401519150606084015160001a90509193909250565b6000816040516020016106a59190611289565b604051602081830303815290604052805190602001209050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106ed826106c2565b9050919050565b6106fd816106e2565b82525050565b600060208201905061071860008301846106f4565b92915050565b6000819050919050565b6107318161071e565b82525050565b600060208201905061074c6000830184610728565b92915050565b6000604051905090565b600080fd5b600080fd5b6000610771826106c2565b9050919050565b61078181610766565b811461078c57600080fd5b50565b60008135905061079e81610778565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6107f7826107ae565b810181811067ffffffffffffffff82111715610816576108156107bf565b5b80604052505050565b6000610829610752565b905061083582826107ee565b919050565b600067ffffffffffffffff821115610855576108546107bf565b5b61085e826107ae565b9050602081019050919050565b82818337600083830152505050565b600061088d6108888461083a565b61081f565b9050828152602081018484840111156108a9576108a86107a9565b5b6108b484828561086b565b509392505050565b600082601f8301126108d1576108d06107a4565b5b81356108e184826020860161087a565b91505092915050565b6108f38161071e565b81146108fe57600080fd5b50565b600081359050610910816108ea565b92915050565b600067ffffffffffffffff821115610931576109306107bf565b5b61093a826107ae565b9050602081019050919050565b600061095a61095584610916565b61081f565b905082815260208101848484011115610976576109756107a9565b5b61098184828561086b565b509392505050565b600082601f83011261099e5761099d6107a4565b5b81356109ae848260208601610947565b91505092915050565b600080600080600060a086880312156109d3576109d261075c565b5b60006109e18882890161078f565b955050602086013567ffffffffffffffff811115610a0257610a01610761565b5b610a0e888289016108bc565b9450506040610a1f88828901610901565b9350506060610a3088828901610901565b925050608086013567ffffffffffffffff811115610a5157610a50610761565b5b610a5d88828901610989565b9150509295509295909350565b6000819050919050565b610a7d81610a6a565b8114610a8857600080fd5b50565b600081359050610a9a81610a74565b92915050565b60008060408385031215610ab757610ab661075c565b5b6000610ac585828601610a8b565b925050602083013567ffffffffffffffff811115610ae657610ae5610761565b5b610af285828601610989565b9150509250929050565b610b05816106e2565b8114610b1057600080fd5b50565b600081359050610b2281610afc565b92915050565b60008060008060808587031215610b4257610b4161075c565b5b6000610b5087828801610b13565b945050602085013567ffffffffffffffff811115610b7157610b70610761565b5b610b7d878288016108bc565b9350506040610b8e87828801610901565b9250506060610b9f87828801610901565b91505092959194509250565b610bb481610a6a565b82525050565b6000602082019050610bcf6000830184610bab565b92915050565b600060208284031215610beb57610bea61075c565b5b600082013567ffffffffffffffff811115610c0957610c08610761565b5b610c1584828501610989565b91505092915050565b600060ff82169050919050565b610c3481610c1e565b82525050565b6000606082019050610c4f6000830186610bab565b610c5c6020830185610bab565b610c696040830184610c2b565b949350505050565b600060208284031215610c8757610c8661075c565b5b6000610c9584828501610a8b565b91505092915050565b600082825260208201905092915050565b7f57726f6e67206e6574776f726b49442121210000000000000000000000000000600082015250565b6000610ce5601283610c9e565b9150610cf082610caf565b602082019050919050565b60006020820190508181036000830152610d1481610cd8565b9050919050565b600081519050919050565b600081905092915050565b60005b83811015610d4f578082015181840152602081019050610d34565b60008484015250505050565b6000610d6682610d1b565b610d708185610d26565b9350610d80818560208601610d31565b80840191505092915050565b6000610d988284610d5b565b915081905092915050565b7f7369676e6174757265206578706972652074696d652121210000000000000000600082015250565b6000610dd9601883610c9e565b9150610de482610da3565b602082019050919050565b60006020820190508181036000830152610e0881610dcc565b9050919050565b7f77726f6e67207265636569766572212121000000000000000000000000000000600082015250565b6000610e45601183610c9e565b9150610e5082610e0f565b602082019050919050565b60006020820190508181036000830152610e7481610e38565b9050919050565b7f636f6e7472616374206973206e6f7420656e6f756768206d6f6e657921212100600082015250565b6000610eb1601f83610c9e565b9150610ebc82610e7b565b602082019050919050565b60006020820190508181036000830152610ee081610ea4565b9050919050565b7f696e76616c6964207369676e6174757265212121000000000000000000000000600082015250565b6000610f1d601483610c9e565b9150610f2882610ee7565b602082019050919050565b60006020820190508181036000830152610f4c81610f10565b9050919050565b6000819050919050565b6000610f78610f73610f6e846106c2565b610f53565b6106c2565b9050919050565b6000610f8a82610f5d565b9050919050565b6000610f9c82610f7f565b9050919050565b610fac81610f91565b82525050565b600082825260208201905092915050565b6000610fce82610d1b565b610fd88185610fb2565b9350610fe8818560208601610d31565b610ff1816107ae565b840191505092915050565b60006060820190506110116000830186610fa3565b81810360208301526110238185610fc3565b90506110326040830184610728565b949350505050565b600060608201905061104f60008301866106f4565b61105c6020830185610728565b6110696040830184610728565b949350505050565b60006080820190506110866000830187610bab565b6110936020830186610c2b565b6110a06040830185610bab565b6110ad6060830184610bab565b95945050505050565b60008160601b9050919050565b60006110ce826110b6565b9050919050565b60006110e0826110c3565b9050919050565b6110f86110f3826106e2565b6110d5565b82525050565b600081519050919050565b600081905092915050565b600061111f826110fe565b6111298185611109565b9350611139818560208601610d31565b80840191505092915050565b6000819050919050565b61116061115b8261071e565b611145565b82525050565b600061117282876110e7565b6014820191506111828286611114565b915061118e828561114f565b60208201915061119e828461114f565b60208201915081905095945050505050565b7f696e76616c6964207369676e6174757265206c656e6774682121210000000000600082015250565b60006111e6601b83610c9e565b91506111f1826111b0565b602082019050919050565b60006020820190508181036000830152611215816111d9565b9050919050565b7f4261746d616e207673204a6f6b65720000000000000000000000000000000000600082015250565b6000611252600f83611109565b915061125d8261121c565b600f82019050919050565b6000819050919050565b61128361127e82610a6a565b611268565b82525050565b600061129482611245565b91506112a08284611272565b6020820191508190509291505056fea26469706673582212202726c3f0a15518cb03b5156030df0b89b52de068304ca944b7592d3211440f5b64736f6c63430008110033",
}

// ExchangecoinABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangecoinMetaData.ABI instead.
var ExchangecoinABI = ExchangecoinMetaData.ABI

// ExchangecoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExchangecoinMetaData.Bin instead.
var ExchangecoinBin = ExchangecoinMetaData.Bin

// DeployExchangecoin deploys a new Ethereum contract, binding an instance of Exchangecoin to it.
func DeployExchangecoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Exchangecoin, error) {
	parsed, err := ExchangecoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExchangecoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Exchangecoin{ExchangecoinCaller: ExchangecoinCaller{contract: contract}, ExchangecoinTransactor: ExchangecoinTransactor{contract: contract}, ExchangecoinFilterer: ExchangecoinFilterer{contract: contract}}, nil
}

// Exchangecoin is an auto generated Go binding around an Ethereum contract.
type Exchangecoin struct {
	ExchangecoinCaller     // Read-only binding to the contract
	ExchangecoinTransactor // Write-only binding to the contract
	ExchangecoinFilterer   // Log filterer for contract events
}

// ExchangecoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangecoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangecoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangecoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangecoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangecoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangecoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangecoinSession struct {
	Contract     *Exchangecoin     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangecoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangecoinCallerSession struct {
	Contract *ExchangecoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ExchangecoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangecoinTransactorSession struct {
	Contract     *ExchangecoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExchangecoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangecoinRaw struct {
	Contract *Exchangecoin // Generic contract binding to access the raw methods on
}

// ExchangecoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangecoinCallerRaw struct {
	Contract *ExchangecoinCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangecoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangecoinTransactorRaw struct {
	Contract *ExchangecoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangecoin creates a new instance of Exchangecoin, bound to a specific deployed contract.
func NewExchangecoin(address common.Address, backend bind.ContractBackend) (*Exchangecoin, error) {
	contract, err := bindExchangecoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchangecoin{ExchangecoinCaller: ExchangecoinCaller{contract: contract}, ExchangecoinTransactor: ExchangecoinTransactor{contract: contract}, ExchangecoinFilterer: ExchangecoinFilterer{contract: contract}}, nil
}

// NewExchangecoinCaller creates a new read-only instance of Exchangecoin, bound to a specific deployed contract.
func NewExchangecoinCaller(address common.Address, caller bind.ContractCaller) (*ExchangecoinCaller, error) {
	contract, err := bindExchangecoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangecoinCaller{contract: contract}, nil
}

// NewExchangecoinTransactor creates a new write-only instance of Exchangecoin, bound to a specific deployed contract.
func NewExchangecoinTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangecoinTransactor, error) {
	contract, err := bindExchangecoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangecoinTransactor{contract: contract}, nil
}

// NewExchangecoinFilterer creates a new log filterer instance of Exchangecoin, bound to a specific deployed contract.
func NewExchangecoinFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangecoinFilterer, error) {
	contract, err := bindExchangecoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangecoinFilterer{contract: contract}, nil
}

// bindExchangecoin binds a generic wrapper to an already deployed contract.
func bindExchangecoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangecoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchangecoin *ExchangecoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchangecoin.Contract.ExchangecoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchangecoin *ExchangecoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchangecoin.Contract.ExchangecoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchangecoin *ExchangecoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchangecoin.Contract.ExchangecoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchangecoin *ExchangecoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchangecoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchangecoin *ExchangecoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchangecoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchangecoin *ExchangecoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchangecoin.Contract.contract.Transact(opts, method, params...)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Exchangecoin *ExchangecoinCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Exchangecoin *ExchangecoinSession) Creator() (common.Address, error) {
	return _Exchangecoin.Contract.Creator(&_Exchangecoin.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Exchangecoin *ExchangecoinCallerSession) Creator() (common.Address, error) {
	return _Exchangecoin.Contract.Creator(&_Exchangecoin.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Exchangecoin *ExchangecoinCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Exchangecoin *ExchangecoinSession) GetBalance() (*big.Int, error) {
	return _Exchangecoin.Contract.GetBalance(&_Exchangecoin.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Exchangecoin *ExchangecoinCallerSession) GetBalance() (*big.Int, error) {
	return _Exchangecoin.Contract.GetBalance(&_Exchangecoin.CallOpts)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 messageHash) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCaller) GetEthSignedMessageHash(opts *bind.CallOpts, messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "getEthSignedMessageHash", messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 messageHash) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinSession) GetEthSignedMessageHash(messageHash [32]byte) ([32]byte, error) {
	return _Exchangecoin.Contract.GetEthSignedMessageHash(&_Exchangecoin.CallOpts, messageHash)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 messageHash) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCallerSession) GetEthSignedMessageHash(messageHash [32]byte) ([32]byte, error) {
	return _Exchangecoin.Contract.GetEthSignedMessageHash(&_Exchangecoin.CallOpts, messageHash)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xa5810c6e.
//
// Solidity: function getMessageHash(address receiver, string message, uint256 amount, uint256 network) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCaller) GetMessageHash(opts *bind.CallOpts, receiver common.Address, message string, amount *big.Int, network *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "getMessageHash", receiver, message, amount, network)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0xa5810c6e.
//
// Solidity: function getMessageHash(address receiver, string message, uint256 amount, uint256 network) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinSession) GetMessageHash(receiver common.Address, message string, amount *big.Int, network *big.Int) ([32]byte, error) {
	return _Exchangecoin.Contract.GetMessageHash(&_Exchangecoin.CallOpts, receiver, message, amount, network)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xa5810c6e.
//
// Solidity: function getMessageHash(address receiver, string message, uint256 amount, uint256 network) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCallerSession) GetMessageHash(receiver common.Address, message string, amount *big.Int, network *big.Int) ([32]byte, error) {
	return _Exchangecoin.Contract.GetMessageHash(&_Exchangecoin.CallOpts, receiver, message, amount, network)
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(uint256)
func (_Exchangecoin *ExchangecoinCaller) NetworkId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "networkId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(uint256)
func (_Exchangecoin *ExchangecoinSession) NetworkId() (*big.Int, error) {
	return _Exchangecoin.Contract.NetworkId(&_Exchangecoin.CallOpts)
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(uint256)
func (_Exchangecoin *ExchangecoinCallerSession) NetworkId() (*big.Int, error) {
	return _Exchangecoin.Contract.NetworkId(&_Exchangecoin.CallOpts)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 ethSignedMessageHash, bytes signature) pure returns(address)
func (_Exchangecoin *ExchangecoinCaller) RecoverSigner(opts *bind.CallOpts, ethSignedMessageHash [32]byte, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "recoverSigner", ethSignedMessageHash, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 ethSignedMessageHash, bytes signature) pure returns(address)
func (_Exchangecoin *ExchangecoinSession) RecoverSigner(ethSignedMessageHash [32]byte, signature []byte) (common.Address, error) {
	return _Exchangecoin.Contract.RecoverSigner(&_Exchangecoin.CallOpts, ethSignedMessageHash, signature)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 ethSignedMessageHash, bytes signature) pure returns(address)
func (_Exchangecoin *ExchangecoinCallerSession) RecoverSigner(ethSignedMessageHash [32]byte, signature []byte) (common.Address, error) {
	return _Exchangecoin.Contract.RecoverSigner(&_Exchangecoin.CallOpts, ethSignedMessageHash, signature)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Exchangecoin *ExchangecoinCaller) SplitSignature(opts *bind.CallOpts, sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "splitSignature", sig)

	outstruct := new(struct {
		R [32]byte
		S [32]byte
		V uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.V = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Exchangecoin *ExchangecoinSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _Exchangecoin.Contract.SplitSignature(&_Exchangecoin.CallOpts, sig)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Exchangecoin *ExchangecoinCallerSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _Exchangecoin.Contract.SplitSignature(&_Exchangecoin.CallOpts, sig)
}

// ReceiveMoney is a paid mutator transaction binding the contract method 0x6d26ec18.
//
// Solidity: function receiveMoney() payable returns()
func (_Exchangecoin *ExchangecoinTransactor) ReceiveMoney(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchangecoin.contract.Transact(opts, "receiveMoney")
}

// ReceiveMoney is a paid mutator transaction binding the contract method 0x6d26ec18.
//
// Solidity: function receiveMoney() payable returns()
func (_Exchangecoin *ExchangecoinSession) ReceiveMoney() (*types.Transaction, error) {
	return _Exchangecoin.Contract.ReceiveMoney(&_Exchangecoin.TransactOpts)
}

// ReceiveMoney is a paid mutator transaction binding the contract method 0x6d26ec18.
//
// Solidity: function receiveMoney() payable returns()
func (_Exchangecoin *ExchangecoinTransactorSession) ReceiveMoney() (*types.Transaction, error) {
	return _Exchangecoin.Contract.ReceiveMoney(&_Exchangecoin.TransactOpts)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0x6c9ee8d1.
//
// Solidity: function withdrawMoney(address receiver, string message, uint256 amount, uint256 network, bytes signature) returns()
func (_Exchangecoin *ExchangecoinTransactor) WithdrawMoney(opts *bind.TransactOpts, receiver common.Address, message string, amount *big.Int, network *big.Int, signature []byte) (*types.Transaction, error) {
	return _Exchangecoin.contract.Transact(opts, "withdrawMoney", receiver, message, amount, network, signature)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0x6c9ee8d1.
//
// Solidity: function withdrawMoney(address receiver, string message, uint256 amount, uint256 network, bytes signature) returns()
func (_Exchangecoin *ExchangecoinSession) WithdrawMoney(receiver common.Address, message string, amount *big.Int, network *big.Int, signature []byte) (*types.Transaction, error) {
	return _Exchangecoin.Contract.WithdrawMoney(&_Exchangecoin.TransactOpts, receiver, message, amount, network, signature)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0x6c9ee8d1.
//
// Solidity: function withdrawMoney(address receiver, string message, uint256 amount, uint256 network, bytes signature) returns()
func (_Exchangecoin *ExchangecoinTransactorSession) WithdrawMoney(receiver common.Address, message string, amount *big.Int, network *big.Int, signature []byte) (*types.Transaction, error) {
	return _Exchangecoin.Contract.WithdrawMoney(&_Exchangecoin.TransactOpts, receiver, message, amount, network, signature)
}

// ExchangecoinEventReceiverMoneyIterator is returned from FilterEventReceiverMoney and is used to iterate over the raw logs and unpacked data for EventReceiverMoney events raised by the Exchangecoin contract.
type ExchangecoinEventReceiverMoneyIterator struct {
	Event *ExchangecoinEventReceiverMoney // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangecoinEventReceiverMoneyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangecoinEventReceiverMoney)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangecoinEventReceiverMoney)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangecoinEventReceiverMoneyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangecoinEventReceiverMoneyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangecoinEventReceiverMoney represents a EventReceiverMoney event raised by the Exchangecoin contract.
type ExchangecoinEventReceiverMoney struct {
	Addr    common.Address
	Amount  *big.Int
	Network *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEventReceiverMoney is a free log retrieval operation binding the contract event 0x2ae6f58e7586dcc615a487ca18d0ed0e5721d25d6920b0a12c30f101ac68c6df.
//
// Solidity: event eventReceiverMoney(address addr, uint256 amount, uint256 network)
func (_Exchangecoin *ExchangecoinFilterer) FilterEventReceiverMoney(opts *bind.FilterOpts) (*ExchangecoinEventReceiverMoneyIterator, error) {

	logs, sub, err := _Exchangecoin.contract.FilterLogs(opts, "eventReceiverMoney")
	if err != nil {
		return nil, err
	}
	return &ExchangecoinEventReceiverMoneyIterator{contract: _Exchangecoin.contract, event: "eventReceiverMoney", logs: logs, sub: sub}, nil
}

// WatchEventReceiverMoney is a free log subscription operation binding the contract event 0x2ae6f58e7586dcc615a487ca18d0ed0e5721d25d6920b0a12c30f101ac68c6df.
//
// Solidity: event eventReceiverMoney(address addr, uint256 amount, uint256 network)
func (_Exchangecoin *ExchangecoinFilterer) WatchEventReceiverMoney(opts *bind.WatchOpts, sink chan<- *ExchangecoinEventReceiverMoney) (event.Subscription, error) {

	logs, sub, err := _Exchangecoin.contract.WatchLogs(opts, "eventReceiverMoney")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangecoinEventReceiverMoney)
				if err := _Exchangecoin.contract.UnpackLog(event, "eventReceiverMoney", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventReceiverMoney is a log parse operation binding the contract event 0x2ae6f58e7586dcc615a487ca18d0ed0e5721d25d6920b0a12c30f101ac68c6df.
//
// Solidity: event eventReceiverMoney(address addr, uint256 amount, uint256 network)
func (_Exchangecoin *ExchangecoinFilterer) ParseEventReceiverMoney(log types.Log) (*ExchangecoinEventReceiverMoney, error) {
	event := new(ExchangecoinEventReceiverMoney)
	if err := _Exchangecoin.contract.UnpackLog(event, "eventReceiverMoney", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangecoinEventWithDrawMoneyIterator is returned from FilterEventWithDrawMoney and is used to iterate over the raw logs and unpacked data for EventWithDrawMoney events raised by the Exchangecoin contract.
type ExchangecoinEventWithDrawMoneyIterator struct {
	Event *ExchangecoinEventWithDrawMoney // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangecoinEventWithDrawMoneyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangecoinEventWithDrawMoney)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangecoinEventWithDrawMoney)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangecoinEventWithDrawMoneyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangecoinEventWithDrawMoneyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangecoinEventWithDrawMoney represents a EventWithDrawMoney event raised by the Exchangecoin contract.
type ExchangecoinEventWithDrawMoney struct {
	Addr      common.Address
	Signature []byte
	Network   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEventWithDrawMoney is a free log retrieval operation binding the contract event 0x33ba68999787348f9a24df6205e34d39f59cbed626f43d7393a29b88706ae949.
//
// Solidity: event eventWithDrawMoney(address addr, bytes signature, uint256 network)
func (_Exchangecoin *ExchangecoinFilterer) FilterEventWithDrawMoney(opts *bind.FilterOpts) (*ExchangecoinEventWithDrawMoneyIterator, error) {

	logs, sub, err := _Exchangecoin.contract.FilterLogs(opts, "eventWithDrawMoney")
	if err != nil {
		return nil, err
	}
	return &ExchangecoinEventWithDrawMoneyIterator{contract: _Exchangecoin.contract, event: "eventWithDrawMoney", logs: logs, sub: sub}, nil
}

// WatchEventWithDrawMoney is a free log subscription operation binding the contract event 0x33ba68999787348f9a24df6205e34d39f59cbed626f43d7393a29b88706ae949.
//
// Solidity: event eventWithDrawMoney(address addr, bytes signature, uint256 network)
func (_Exchangecoin *ExchangecoinFilterer) WatchEventWithDrawMoney(opts *bind.WatchOpts, sink chan<- *ExchangecoinEventWithDrawMoney) (event.Subscription, error) {

	logs, sub, err := _Exchangecoin.contract.WatchLogs(opts, "eventWithDrawMoney")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangecoinEventWithDrawMoney)
				if err := _Exchangecoin.contract.UnpackLog(event, "eventWithDrawMoney", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventWithDrawMoney is a log parse operation binding the contract event 0x33ba68999787348f9a24df6205e34d39f59cbed626f43d7393a29b88706ae949.
//
// Solidity: event eventWithDrawMoney(address addr, bytes signature, uint256 network)
func (_Exchangecoin *ExchangecoinFilterer) ParseEventWithDrawMoney(log types.Log) (*ExchangecoinEventWithDrawMoney, error) {
	event := new(ExchangecoinEventWithDrawMoney)
	if err := _Exchangecoin.contract.UnpackLog(event, "eventWithDrawMoney", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
