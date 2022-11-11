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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveMoney\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdrawMoney\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610cba806100606000396000f3fe60806040526004361061003f5760003560e01c806302d05d3f1461004457806312065fe01461006f5780636c9ee8d11461009a5780636d26ec18146100d7575b600080fd5b34801561005057600080fd5b506100596100e1565b60405161006691906104c1565b60405180910390f35b34801561007b57600080fd5b50610084610105565b60405161009191906104f5565b60405180910390f35b3480156100a657600080fd5b506100c160048036038101906100bc9190610763565b61010d565b6040516100ce9190610831565b60405180910390f35b6100df6102b0565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600047905090565b60008061011c87878787610340565b9050600061012982610379565b905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661016c82866103a9565b73ffffffffffffffffffffffffffffffffffffffff16146101c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b9906108a9565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614610230576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022790610915565b60405180910390fd5b60003373ffffffffffffffffffffffffffffffffffffffff16876113889060405161025a90610966565b600060405180830381858888f193505050503d8060008114610298576040519150601f19603f3d011682016040523d82523d6000602084013e61029d565b606091505b5050905080935050505095945050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461033e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610335906109c7565b60405180910390fd5b565b6000848484846040516020016103599493929190610ac1565b604051602081830303815290604052805190602001209050949350505050565b60008160405160200161038c9190610b82565b604051602081830303815290604052805190602001209050919050565b6000806000806103b885610418565b925092509250600186828585604051600081526020016040526040516103e19493929190610bd3565b6020604051602081039080840390855afa158015610403573d6000803e3d6000fd5b50505060206040510351935050505092915050565b60008060006041845114610461576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045890610c64565b60405180910390fd5b6020840151925060408401519150606084015160001a90509193909250565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104ab82610480565b9050919050565b6104bb816104a0565b82525050565b60006020820190506104d660008301846104b2565b92915050565b6000819050919050565b6104ef816104dc565b82525050565b600060208201905061050a60008301846104e6565b92915050565b6000604051905090565b600080fd5b600080fd5b61052d816104a0565b811461053857600080fd5b50565b60008135905061054a81610524565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6105a38261055a565b810181811067ffffffffffffffff821117156105c2576105c161056b565b5b80604052505050565b60006105d5610510565b90506105e1828261059a565b919050565b600067ffffffffffffffff8211156106015761060061056b565b5b61060a8261055a565b9050602081019050919050565b82818337600083830152505050565b6000610639610634846105e6565b6105cb565b90508281526020810184848401111561065557610654610555565b5b610660848285610617565b509392505050565b600082601f83011261067d5761067c610550565b5b813561068d848260208601610626565b91505092915050565b61069f816104dc565b81146106aa57600080fd5b50565b6000813590506106bc81610696565b92915050565b600067ffffffffffffffff8211156106dd576106dc61056b565b5b6106e68261055a565b9050602081019050919050565b6000610706610701846106c2565b6105cb565b90508281526020810184848401111561072257610721610555565b5b61072d848285610617565b509392505050565b600082601f83011261074a57610749610550565b5b813561075a8482602086016106f3565b91505092915050565b600080600080600060a0868803121561077f5761077e61051a565b5b600061078d8882890161053b565b955050602086013567ffffffffffffffff8111156107ae576107ad61051f565b5b6107ba88828901610668565b94505060406107cb888289016106ad565b93505060606107dc888289016106ad565b925050608086013567ffffffffffffffff8111156107fd576107fc61051f565b5b61080988828901610735565b9150509295509295909350565b60008115159050919050565b61082b81610816565b82525050565b60006020820190506108466000830184610822565b92915050565b600082825260208201905092915050565b7f696e76616c6964207369676e6174757265212121000000000000000000000000600082015250565b600061089360148361084c565b915061089e8261085d565b602082019050919050565b600060208201905081810360008301526108c281610886565b9050919050565b7f77726f6e67207265636569766572212121000000000000000000000000000000600082015250565b60006108ff60118361084c565b915061090a826108c9565b602082019050919050565b6000602082019050818103600083015261092e816108f2565b9050919050565b600081905092915050565b50565b6000610950600083610935565b915061095b82610940565b600082019050919050565b600061097182610943565b9150819050919050565b7f4e6f742061757468656e7469636174696f6e2121210000000000000000000000600082015250565b60006109b160158361084c565b91506109bc8261097b565b602082019050919050565b600060208201905081810360008301526109e0816109a4565b9050919050565b60008160601b9050919050565b60006109ff826109e7565b9050919050565b6000610a11826109f4565b9050919050565b610a29610a24826104a0565b610a06565b82525050565b600081519050919050565b600081905092915050565b60005b83811015610a63578082015181840152602081019050610a48565b60008484015250505050565b6000610a7a82610a2f565b610a848185610a3a565b9350610a94818560208601610a45565b80840191505092915050565b6000819050919050565b610abb610ab6826104dc565b610aa0565b82525050565b6000610acd8287610a18565b601482019150610add8286610a6f565b9150610ae98285610aaa565b602082019150610af98284610aaa565b60208201915081905095945050505050565b7f4261746d616e207673204a6f6b65720000000000000000000000000000000000600082015250565b6000610b41600f83610a3a565b9150610b4c82610b0b565b600f82019050919050565b6000819050919050565b6000819050919050565b610b7c610b7782610b57565b610b61565b82525050565b6000610b8d82610b34565b9150610b998284610b6b565b60208201915081905092915050565b610bb181610b57565b82525050565b600060ff82169050919050565b610bcd81610bb7565b82525050565b6000608082019050610be86000830187610ba8565b610bf56020830186610bc4565b610c026040830185610ba8565b610c0f6060830184610ba8565b95945050505050565b7f696e76616c6964207369676e6174757265206c656e6774682121210000000000600082015250565b6000610c4e601b8361084c565b9150610c5982610c18565b602082019050919050565b60006020820190508181036000830152610c7d81610c41565b905091905056fea2646970667358221220be611fa26818d5cd26859d61df364d275277c0aeb1688ffabcb087fdb79cefc264736f6c63430008110033",
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
