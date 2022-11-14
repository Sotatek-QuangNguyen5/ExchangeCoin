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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"onSetNumber\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveMoney\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdrawMoney\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610e97806100606000396000f3fe6080604052600436106100705760003560e01c80633fb5c1cb1161004e5780633fb5c1cb146100f65780636c9ee8d11461011f5780636d26ec181461015c578063f2c9ecd81461016657610070565b806302d05d3f1461007557806312065fe0146100a057806323fd0e40146100cb575b600080fd5b34801561008157600080fd5b5061008a610191565b6040516100979190610605565b60405180910390f35b3480156100ac57600080fd5b506100b56101b5565b6040516100c29190610639565b60405180910390f35b3480156100d757600080fd5b506100e06101bd565b6040516100ed9190610639565b60405180910390f35b34801561010257600080fd5b5061011d60048036038101906101189190610694565b6101c3565b005b34801561012b57600080fd5b50610146600480360381019061014191906108d4565b610204565b60405161015391906109a2565b60405180910390f35b6101646103ea565b005b34801561017257600080fd5b5061017b61047a565b6040516101889190610639565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600047905090565b60015481565b7f1f6641444d11c93a5ca4da333655135a1b21c94ed0a6ba2f156a0dfc4626a168336040516101f29190610605565b60405180910390a18060018190555050565b60008061021387878787610484565b90506000610220826104bd565b905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661026382866104ed565b73ffffffffffffffffffffffffffffffffffffffff16146102b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b090610a1a565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614610327576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031e90610a86565b60405180910390fd5b4786111561036a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036190610af2565b60405180910390fd5b60003373ffffffffffffffffffffffffffffffffffffffff16876113889060405161039490610b43565b600060405180830381858888f193505050503d80600081146103d2576040519150601f19603f3d011682016040523d82523d6000602084013e6103d7565b606091505b5050905080935050505095945050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610478576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046f90610ba4565b60405180910390fd5b565b6000600154905090565b60008484848460405160200161049d9493929190610c9e565b604051602081830303815290604052805190602001209050949350505050565b6000816040516020016104d09190610d5f565b604051602081830303815290604052805190602001209050919050565b6000806000806104fc8561055c565b925092509250600186828585604051600081526020016040526040516105259493929190610db0565b6020604051602081039080840390855afa158015610547573d6000803e3d6000fd5b50505060206040510351935050505092915050565b600080600060418451146105a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059c90610e41565b60405180910390fd5b6020840151925060408401519150606084015160001a90509193909250565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105ef826105c4565b9050919050565b6105ff816105e4565b82525050565b600060208201905061061a60008301846105f6565b92915050565b6000819050919050565b61063381610620565b82525050565b600060208201905061064e600083018461062a565b92915050565b6000604051905090565b600080fd5b600080fd5b61067181610620565b811461067c57600080fd5b50565b60008135905061068e81610668565b92915050565b6000602082840312156106aa576106a961065e565b5b60006106b88482850161067f565b91505092915050565b6106ca816105e4565b81146106d557600080fd5b50565b6000813590506106e7816106c1565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610740826106f7565b810181811067ffffffffffffffff8211171561075f5761075e610708565b5b80604052505050565b6000610772610654565b905061077e8282610737565b919050565b600067ffffffffffffffff82111561079e5761079d610708565b5b6107a7826106f7565b9050602081019050919050565b82818337600083830152505050565b60006107d66107d184610783565b610768565b9050828152602081018484840111156107f2576107f16106f2565b5b6107fd8482856107b4565b509392505050565b600082601f83011261081a576108196106ed565b5b813561082a8482602086016107c3565b91505092915050565b600067ffffffffffffffff82111561084e5761084d610708565b5b610857826106f7565b9050602081019050919050565b600061087761087284610833565b610768565b905082815260208101848484011115610893576108926106f2565b5b61089e8482856107b4565b509392505050565b600082601f8301126108bb576108ba6106ed565b5b81356108cb848260208601610864565b91505092915050565b600080600080600060a086880312156108f0576108ef61065e565b5b60006108fe888289016106d8565b955050602086013567ffffffffffffffff81111561091f5761091e610663565b5b61092b88828901610805565b945050604061093c8882890161067f565b935050606061094d8882890161067f565b925050608086013567ffffffffffffffff81111561096e5761096d610663565b5b61097a888289016108a6565b9150509295509295909350565b60008115159050919050565b61099c81610987565b82525050565b60006020820190506109b76000830184610993565b92915050565b600082825260208201905092915050565b7f696e76616c6964207369676e6174757265212121000000000000000000000000600082015250565b6000610a046014836109bd565b9150610a0f826109ce565b602082019050919050565b60006020820190508181036000830152610a33816109f7565b9050919050565b7f77726f6e67207265636569766572212121000000000000000000000000000000600082015250565b6000610a706011836109bd565b9150610a7b82610a3a565b602082019050919050565b60006020820190508181036000830152610a9f81610a63565b9050919050565b7f6e6f7420656e6f756768206d6f6e657921212100000000000000000000000000600082015250565b6000610adc6013836109bd565b9150610ae782610aa6565b602082019050919050565b60006020820190508181036000830152610b0b81610acf565b9050919050565b600081905092915050565b50565b6000610b2d600083610b12565b9150610b3882610b1d565b600082019050919050565b6000610b4e82610b20565b9150819050919050565b7f4e6f742061757468656e7469636174696f6e2121210000000000000000000000600082015250565b6000610b8e6015836109bd565b9150610b9982610b58565b602082019050919050565b60006020820190508181036000830152610bbd81610b81565b9050919050565b60008160601b9050919050565b6000610bdc82610bc4565b9050919050565b6000610bee82610bd1565b9050919050565b610c06610c01826105e4565b610be3565b82525050565b600081519050919050565b600081905092915050565b60005b83811015610c40578082015181840152602081019050610c25565b60008484015250505050565b6000610c5782610c0c565b610c618185610c17565b9350610c71818560208601610c22565b80840191505092915050565b6000819050919050565b610c98610c9382610620565b610c7d565b82525050565b6000610caa8287610bf5565b601482019150610cba8286610c4c565b9150610cc68285610c87565b602082019150610cd68284610c87565b60208201915081905095945050505050565b7f4261746d616e207673204a6f6b65720000000000000000000000000000000000600082015250565b6000610d1e600f83610c17565b9150610d2982610ce8565b600f82019050919050565b6000819050919050565b6000819050919050565b610d59610d5482610d34565b610d3e565b82525050565b6000610d6a82610d11565b9150610d768284610d48565b60208201915081905092915050565b610d8e81610d34565b82525050565b600060ff82169050919050565b610daa81610d94565b82525050565b6000608082019050610dc56000830187610d85565b610dd26020830186610da1565b610ddf6040830185610d85565b610dec6060830184610d85565b95945050505050565b7f696e76616c6964207369676e6174757265206c656e6774682121210000000000600082015250565b6000610e2b601b836109bd565b9150610e3682610df5565b602082019050919050565b60006020820190508181036000830152610e5a81610e1e565b905091905056fea2646970667358221220f7759672a873d90d1bad0ce825eaaed613b5dadc2de93f9606315997e48e744d64736f6c63430008110033",
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

// GetNumber is a free data retrieval call binding the contract method 0xf2c9ecd8.
//
// Solidity: function getNumber() view returns(uint256)
func (_Exchangecoin *ExchangecoinCaller) GetNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "getNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumber is a free data retrieval call binding the contract method 0xf2c9ecd8.
//
// Solidity: function getNumber() view returns(uint256)
func (_Exchangecoin *ExchangecoinSession) GetNumber() (*big.Int, error) {
	return _Exchangecoin.Contract.GetNumber(&_Exchangecoin.CallOpts)
}

// GetNumber is a free data retrieval call binding the contract method 0xf2c9ecd8.
//
// Solidity: function getNumber() view returns(uint256)
func (_Exchangecoin *ExchangecoinCallerSession) GetNumber() (*big.Int, error) {
	return _Exchangecoin.Contract.GetNumber(&_Exchangecoin.CallOpts)
}

// MyNumber is a free data retrieval call binding the contract method 0x23fd0e40.
//
// Solidity: function myNumber() view returns(uint256)
func (_Exchangecoin *ExchangecoinCaller) MyNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "myNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyNumber is a free data retrieval call binding the contract method 0x23fd0e40.
//
// Solidity: function myNumber() view returns(uint256)
func (_Exchangecoin *ExchangecoinSession) MyNumber() (*big.Int, error) {
	return _Exchangecoin.Contract.MyNumber(&_Exchangecoin.CallOpts)
}

// MyNumber is a free data retrieval call binding the contract method 0x23fd0e40.
//
// Solidity: function myNumber() view returns(uint256)
func (_Exchangecoin *ExchangecoinCallerSession) MyNumber() (*big.Int, error) {
	return _Exchangecoin.Contract.MyNumber(&_Exchangecoin.CallOpts)
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

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 number) returns()
func (_Exchangecoin *ExchangecoinTransactor) SetNumber(opts *bind.TransactOpts, number *big.Int) (*types.Transaction, error) {
	return _Exchangecoin.contract.Transact(opts, "setNumber", number)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 number) returns()
func (_Exchangecoin *ExchangecoinSession) SetNumber(number *big.Int) (*types.Transaction, error) {
	return _Exchangecoin.Contract.SetNumber(&_Exchangecoin.TransactOpts, number)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 number) returns()
func (_Exchangecoin *ExchangecoinTransactorSession) SetNumber(number *big.Int) (*types.Transaction, error) {
	return _Exchangecoin.Contract.SetNumber(&_Exchangecoin.TransactOpts, number)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0x6c9ee8d1.
//
// Solidity: function withdrawMoney(address receiver, string message, uint256 amount, uint256 nonce, bytes signature) returns(bool)
func (_Exchangecoin *ExchangecoinTransactor) WithdrawMoney(opts *bind.TransactOpts, receiver common.Address, message string, amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Exchangecoin.contract.Transact(opts, "withdrawMoney", receiver, message, amount, nonce, signature)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0x6c9ee8d1.
//
// Solidity: function withdrawMoney(address receiver, string message, uint256 amount, uint256 nonce, bytes signature) returns(bool)
func (_Exchangecoin *ExchangecoinSession) WithdrawMoney(receiver common.Address, message string, amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Exchangecoin.Contract.WithdrawMoney(&_Exchangecoin.TransactOpts, receiver, message, amount, nonce, signature)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0x6c9ee8d1.
//
// Solidity: function withdrawMoney(address receiver, string message, uint256 amount, uint256 nonce, bytes signature) returns(bool)
func (_Exchangecoin *ExchangecoinTransactorSession) WithdrawMoney(receiver common.Address, message string, amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Exchangecoin.Contract.WithdrawMoney(&_Exchangecoin.TransactOpts, receiver, message, amount, nonce, signature)
}

// ExchangecoinOnSetNumberIterator is returned from FilterOnSetNumber and is used to iterate over the raw logs and unpacked data for OnSetNumber events raised by the Exchangecoin contract.
type ExchangecoinOnSetNumberIterator struct {
	Event *ExchangecoinOnSetNumber // Event containing the contract specifics and raw log

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
func (it *ExchangecoinOnSetNumberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangecoinOnSetNumber)
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
		it.Event = new(ExchangecoinOnSetNumber)
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
func (it *ExchangecoinOnSetNumberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangecoinOnSetNumberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangecoinOnSetNumber represents a OnSetNumber event raised by the Exchangecoin contract.
type ExchangecoinOnSetNumber struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnSetNumber is a free log retrieval operation binding the contract event 0x1f6641444d11c93a5ca4da333655135a1b21c94ed0a6ba2f156a0dfc4626a168.
//
// Solidity: event onSetNumber(address addr)
func (_Exchangecoin *ExchangecoinFilterer) FilterOnSetNumber(opts *bind.FilterOpts) (*ExchangecoinOnSetNumberIterator, error) {

	logs, sub, err := _Exchangecoin.contract.FilterLogs(opts, "onSetNumber")
	if err != nil {
		return nil, err
	}
	return &ExchangecoinOnSetNumberIterator{contract: _Exchangecoin.contract, event: "onSetNumber", logs: logs, sub: sub}, nil
}

// WatchOnSetNumber is a free log subscription operation binding the contract event 0x1f6641444d11c93a5ca4da333655135a1b21c94ed0a6ba2f156a0dfc4626a168.
//
// Solidity: event onSetNumber(address addr)
func (_Exchangecoin *ExchangecoinFilterer) WatchOnSetNumber(opts *bind.WatchOpts, sink chan<- *ExchangecoinOnSetNumber) (event.Subscription, error) {

	logs, sub, err := _Exchangecoin.contract.WatchLogs(opts, "onSetNumber")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangecoinOnSetNumber)
				if err := _Exchangecoin.contract.UnpackLog(event, "onSetNumber", log); err != nil {
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

// ParseOnSetNumber is a log parse operation binding the contract event 0x1f6641444d11c93a5ca4da333655135a1b21c94ed0a6ba2f156a0dfc4626a168.
//
// Solidity: event onSetNumber(address addr)
func (_Exchangecoin *ExchangecoinFilterer) ParseOnSetNumber(log types.Log) (*ExchangecoinOnSetNumber, error) {
	event := new(ExchangecoinOnSetNumber)
	if err := _Exchangecoin.contract.UnpackLog(event, "onSetNumber", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
