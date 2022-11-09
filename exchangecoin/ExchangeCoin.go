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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"getEthSignedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethSignedMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610a08806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632dd34f0f1461005c57806397aba7f91461008c578063a7bb5803146100bc578063b446f3b2146100ee578063fa5408011461011e575b600080fd5b6100766004803603810190610071919061053c565b61014e565b60405161008391906105e2565b60405180910390f35b6100a660048036038101906100a19190610633565b6101ac565b6040516100b3919061069e565b60405180910390f35b6100d660048036038101906100d191906106b9565b61021b565b6040516100e59392919061072d565b60405180910390f35b61010860048036038101906101039190610764565b610283565b60405161011591906107ad565b60405180910390f35b610138600480360381019061013391906107c8565b6102b3565b60405161014591906107ad565b60405180910390f35b60008061015a84610283565b90506000610167826102b3565b90508573ffffffffffffffffffffffffffffffffffffffff1661018a82866101ac565b73ffffffffffffffffffffffffffffffffffffffff1614925050509392505050565b6000806000806101bb8561021b565b925092509250600186828585604051600081526020016040526040516101e494939291906107f5565b6020604051602081039080840390855afa158015610206573d6000803e3d6000fd5b50505060206040510351935050505092915050565b60008060006041845114610264576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025b90610897565b60405180910390fd5b6020840151925060408401519150606084015160001a90509193909250565b6000816040516020016102969190610928565b604051602081830303815290604052805190602001209050919050565b6000816040516020016102c691906109ac565b604051602081830303815290604052805190602001209050919050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610322826102f7565b9050919050565b61033281610317565b811461033d57600080fd5b50565b60008135905061034f81610329565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6103a88261035f565b810181811067ffffffffffffffff821117156103c7576103c6610370565b5b80604052505050565b60006103da6102e3565b90506103e6828261039f565b919050565b600067ffffffffffffffff82111561040657610405610370565b5b61040f8261035f565b9050602081019050919050565b82818337600083830152505050565b600061043e610439846103eb565b6103d0565b90508281526020810184848401111561045a5761045961035a565b5b61046584828561041c565b509392505050565b600082601f83011261048257610481610355565b5b813561049284826020860161042b565b91505092915050565b600067ffffffffffffffff8211156104b6576104b5610370565b5b6104bf8261035f565b9050602081019050919050565b60006104df6104da8461049b565b6103d0565b9050828152602081018484840111156104fb576104fa61035a565b5b61050684828561041c565b509392505050565b600082601f83011261052357610522610355565b5b81356105338482602086016104cc565b91505092915050565b600080600060608486031215610555576105546102ed565b5b600061056386828701610340565b935050602084013567ffffffffffffffff811115610584576105836102f2565b5b6105908682870161046d565b925050604084013567ffffffffffffffff8111156105b1576105b06102f2565b5b6105bd8682870161050e565b9150509250925092565b60008115159050919050565b6105dc816105c7565b82525050565b60006020820190506105f760008301846105d3565b92915050565b6000819050919050565b610610816105fd565b811461061b57600080fd5b50565b60008135905061062d81610607565b92915050565b6000806040838503121561064a576106496102ed565b5b60006106588582860161061e565b925050602083013567ffffffffffffffff811115610679576106786102f2565b5b6106858582860161050e565b9150509250929050565b61069881610317565b82525050565b60006020820190506106b3600083018461068f565b92915050565b6000602082840312156106cf576106ce6102ed565b5b600082013567ffffffffffffffff8111156106ed576106ec6102f2565b5b6106f98482850161050e565b91505092915050565b61070b816105fd565b82525050565b600060ff82169050919050565b61072781610711565b82525050565b60006060820190506107426000830186610702565b61074f6020830185610702565b61075c604083018461071e565b949350505050565b60006020828403121561077a576107796102ed565b5b600082013567ffffffffffffffff811115610798576107976102f2565b5b6107a48482850161046d565b91505092915050565b60006020820190506107c26000830184610702565b92915050565b6000602082840312156107de576107dd6102ed565b5b60006107ec8482850161061e565b91505092915050565b600060808201905061080a6000830187610702565b610817602083018661071e565b6108246040830185610702565b6108316060830184610702565b95945050505050565b600082825260208201905092915050565b7f696e76616c6964207369676e6174757265206c656e6774680000000000000000600082015250565b600061088160188361083a565b915061088c8261084b565b602082019050919050565b600060208201905081810360008301526108b081610874565b9050919050565b600081519050919050565b600081905092915050565b60005b838110156108eb5780820151818401526020810190506108d0565b60008484015250505050565b6000610902826108b7565b61090c81856108c2565b935061091c8185602086016108cd565b80840191505092915050565b600061093482846108f7565b915081905092915050565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600082015250565b6000610975601c836108c2565b91506109808261093f565b601c82019050919050565b6000819050919050565b6109a66109a1826105fd565b61098b565b82525050565b60006109b782610968565b91506109c38284610995565b6020820191508190509291505056fea264697066735822122087cbe2b4ad3b4c48569737bd439663f45337f0a5c1d6972717fd91a52380181464736f6c63430008110033",
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

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCaller) GetEthSignedMessageHash(opts *bind.CallOpts, _messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "getEthSignedMessageHash", _messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _Exchangecoin.Contract.GetEthSignedMessageHash(&_Exchangecoin.CallOpts, _messageHash)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCallerSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _Exchangecoin.Contract.GetEthSignedMessageHash(&_Exchangecoin.CallOpts, _messageHash)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCaller) GetMessageHash(opts *bind.CallOpts, _message string) ([32]byte, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "getMessageHash", _message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinSession) GetMessageHash(_message string) ([32]byte, error) {
	return _Exchangecoin.Contract.GetMessageHash(&_Exchangecoin.CallOpts, _message)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCallerSession) GetMessageHash(_message string) ([32]byte, error) {
	return _Exchangecoin.Contract.GetMessageHash(&_Exchangecoin.CallOpts, _message)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Exchangecoin *ExchangecoinCaller) RecoverSigner(opts *bind.CallOpts, _ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "recoverSigner", _ethSignedMessageHash, _signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Exchangecoin *ExchangecoinSession) RecoverSigner(_ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	return _Exchangecoin.Contract.RecoverSigner(&_Exchangecoin.CallOpts, _ethSignedMessageHash, _signature)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Exchangecoin *ExchangecoinCallerSession) RecoverSigner(_ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	return _Exchangecoin.Contract.RecoverSigner(&_Exchangecoin.CallOpts, _ethSignedMessageHash, _signature)
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

// Verify is a free data retrieval call binding the contract method 0x2dd34f0f.
//
// Solidity: function verify(address signer, string _message, bytes signature) pure returns(bool)
func (_Exchangecoin *ExchangecoinCaller) Verify(opts *bind.CallOpts, signer common.Address, _message string, signature []byte) (bool, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "verify", signer, _message, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x2dd34f0f.
//
// Solidity: function verify(address signer, string _message, bytes signature) pure returns(bool)
func (_Exchangecoin *ExchangecoinSession) Verify(signer common.Address, _message string, signature []byte) (bool, error) {
	return _Exchangecoin.Contract.Verify(&_Exchangecoin.CallOpts, signer, _message, signature)
}

// Verify is a free data retrieval call binding the contract method 0x2dd34f0f.
//
// Solidity: function verify(address signer, string _message, bytes signature) pure returns(bool)
func (_Exchangecoin *ExchangecoinCallerSession) Verify(signer common.Address, _message string, signature []byte) (bool, error) {
	return _Exchangecoin.Contract.Verify(&_Exchangecoin.CallOpts, signer, _message, signature)
}
