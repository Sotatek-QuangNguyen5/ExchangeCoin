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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"eventReceiverMoney\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"eventWithDrawMoney\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"getEthSignedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveMoney\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ethSignedMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdrawMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611199806100606000396000f3fe60806040526004361061007b5760003560e01c80636d26ec181161004e5780636d26ec181461013c57806397aba7f914610146578063a7bb580314610183578063fa540801146101c25761007b565b806302d05d3f14610080578063082d8f29146100ab57806312065fe0146100e857806348bf759b14610113575b600080fd5b34801561008c57600080fd5b506100956101ff565b6040516100a29190610678565b60405180910390f35b3480156100b757600080fd5b506100d260048036038101906100cd919061084f565b610223565b6040516100df91906108d7565b60405180910390f35b3480156100f457600080fd5b506100fd610259565b60405161010a9190610901565b60405180910390f35b34801561011f57600080fd5b5061013a600480360381019061013591906109fb565b610261565b005b6101446104f5565b005b34801561015257600080fd5b5061016d60048036038101906101689190610ac6565b610530565b60405161017a9190610678565b60405180910390f35b34801561018f57600080fd5b506101aa60048036038101906101a59190610b22565b61059f565b6040516101b993929190610b87565b60405180910390f35b3480156101ce57600080fd5b506101e960048036038101906101e49190610bbe565b610607565b6040516101f691906108d7565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600083838360405160200161023a93929190610cc5565b6040516020818303038152906040528051906020012090509392505050565b600047905090565b600015156001826040516102759190610d45565b908152602001604051809103902060009054906101000a900460ff161515146102d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102ca90610db9565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614610341576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033890610e25565b60405180910390fd5b47821115610384576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037b90610e91565b60405180910390fd5b6000610391858585610223565b9050600061039e82610607565b905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166103e18285610530565b73ffffffffffffffffffffffffffffffffffffffff1614610437576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042e90610efd565b60405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff166108fc859081150290604051600060405180830381858888f1935050505015801561047d573d6000803e3d6000fd5b506001808460405161048f9190610d45565b908152602001604051809103902060006101000a81548160ff0219169083151502179055507fcf3bb011b58dadd278f376bd3eca311ad5135ee846d6efab8eff90dd0e1ddf0d86846040516104e5929190610fc6565b60405180910390a1505050505050565b7fb0b6cfca2b3037ba69e8256ed0cf4e16b8d93ac32f69acb036c2c49576886ac03334604051610526929190610ff6565b60405180910390a1565b60008060008061053f8561059f565b92509250925060018682858560405160008152602001604052604051610568949392919061101f565b6020604051602081039080840390855afa15801561058a573d6000803e3d6000fd5b50505060206040510351935050505092915050565b600080600060418451146105e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105df906110b0565b60405180910390fd5b6020840151925060408401519150606084015160001a90509193909250565b60008160405160200161061a919061113d565b604051602081830303815290604052805190602001209050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061066282610637565b9050919050565b61067281610657565b82525050565b600060208201905061068d6000830184610669565b92915050565b6000604051905090565b600080fd5b600080fd5b6106b081610657565b81146106bb57600080fd5b50565b6000813590506106cd816106a7565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610726826106dd565b810181811067ffffffffffffffff82111715610745576107446106ee565b5b80604052505050565b6000610758610693565b9050610764828261071d565b919050565b600067ffffffffffffffff821115610784576107836106ee565b5b61078d826106dd565b9050602081019050919050565b82818337600083830152505050565b60006107bc6107b784610769565b61074e565b9050828152602081018484840111156107d8576107d76106d8565b5b6107e384828561079a565b509392505050565b600082601f830112610800576107ff6106d3565b5b81356108108482602086016107a9565b91505092915050565b6000819050919050565b61082c81610819565b811461083757600080fd5b50565b60008135905061084981610823565b92915050565b6000806000606084860312156108685761086761069d565b5b6000610876868287016106be565b935050602084013567ffffffffffffffff811115610897576108966106a2565b5b6108a3868287016107eb565b92505060406108b48682870161083a565b9150509250925092565b6000819050919050565b6108d1816108be565b82525050565b60006020820190506108ec60008301846108c8565b92915050565b6108fb81610819565b82525050565b600060208201905061091660008301846108f2565b92915050565b600061092782610637565b9050919050565b6109378161091c565b811461094257600080fd5b50565b6000813590506109548161092e565b92915050565b600067ffffffffffffffff821115610975576109746106ee565b5b61097e826106dd565b9050602081019050919050565b600061099e6109998461095a565b61074e565b9050828152602081018484840111156109ba576109b96106d8565b5b6109c584828561079a565b509392505050565b600082601f8301126109e2576109e16106d3565b5b81356109f284826020860161098b565b91505092915050565b60008060008060808587031215610a1557610a1461069d565b5b6000610a2387828801610945565b945050602085013567ffffffffffffffff811115610a4457610a436106a2565b5b610a50878288016107eb565b9350506040610a618782880161083a565b925050606085013567ffffffffffffffff811115610a8257610a816106a2565b5b610a8e878288016109cd565b91505092959194509250565b610aa3816108be565b8114610aae57600080fd5b50565b600081359050610ac081610a9a565b92915050565b60008060408385031215610add57610adc61069d565b5b6000610aeb85828601610ab1565b925050602083013567ffffffffffffffff811115610b0c57610b0b6106a2565b5b610b18858286016109cd565b9150509250929050565b600060208284031215610b3857610b3761069d565b5b600082013567ffffffffffffffff811115610b5657610b556106a2565b5b610b62848285016109cd565b91505092915050565b600060ff82169050919050565b610b8181610b6b565b82525050565b6000606082019050610b9c60008301866108c8565b610ba960208301856108c8565b610bb66040830184610b78565b949350505050565b600060208284031215610bd457610bd361069d565b5b6000610be284828501610ab1565b91505092915050565b60008160601b9050919050565b6000610c0382610beb565b9050919050565b6000610c1582610bf8565b9050919050565b610c2d610c2882610657565b610c0a565b82525050565b600081519050919050565b600081905092915050565b60005b83811015610c67578082015181840152602081019050610c4c565b60008484015250505050565b6000610c7e82610c33565b610c888185610c3e565b9350610c98818560208601610c49565b80840191505092915050565b6000819050919050565b610cbf610cba82610819565b610ca4565b82525050565b6000610cd18286610c1c565b601482019150610ce18285610c73565b9150610ced8284610cae565b602082019150819050949350505050565b600081519050919050565b600081905092915050565b6000610d1f82610cfe565b610d298185610d09565b9350610d39818560208601610c49565b80840191505092915050565b6000610d518284610d14565b915081905092915050565b600082825260208201905092915050565b7f7369676e6174757265206578706972652074696d652121210000000000000000600082015250565b6000610da3601883610d5c565b9150610dae82610d6d565b602082019050919050565b60006020820190508181036000830152610dd281610d96565b9050919050565b7f77726f6e67207265636569766572212121000000000000000000000000000000600082015250565b6000610e0f601183610d5c565b9150610e1a82610dd9565b602082019050919050565b60006020820190508181036000830152610e3e81610e02565b9050919050565b7f636f6e7472616374206973206e6f7420656e6f756768206d6f6e657921212100600082015250565b6000610e7b601f83610d5c565b9150610e8682610e45565b602082019050919050565b60006020820190508181036000830152610eaa81610e6e565b9050919050565b7f696e76616c6964207369676e6174757265212121000000000000000000000000600082015250565b6000610ee7601483610d5c565b9150610ef282610eb1565b602082019050919050565b60006020820190508181036000830152610f1681610eda565b9050919050565b6000819050919050565b6000610f42610f3d610f3884610637565b610f1d565b610637565b9050919050565b6000610f5482610f27565b9050919050565b6000610f6682610f49565b9050919050565b610f7681610f5b565b82525050565b600082825260208201905092915050565b6000610f9882610cfe565b610fa28185610f7c565b9350610fb2818560208601610c49565b610fbb816106dd565b840191505092915050565b6000604082019050610fdb6000830185610f6d565b8181036020830152610fed8184610f8d565b90509392505050565b600060408201905061100b6000830185610669565b61101860208301846108f2565b9392505050565b600060808201905061103460008301876108c8565b6110416020830186610b78565b61104e60408301856108c8565b61105b60608301846108c8565b95945050505050565b7f696e76616c6964207369676e6174757265206c656e6774682121210000000000600082015250565b600061109a601b83610d5c565b91506110a582611064565b602082019050919050565b600060208201905081810360008301526110c98161108d565b9050919050565b7f4261746d616e207673204a6f6b65720000000000000000000000000000000000600082015250565b6000611106600f83610c3e565b9150611111826110d0565b600f82019050919050565b6000819050919050565b611137611132826108be565b61111c565b82525050565b6000611148826110f9565b91506111548284611126565b6020820191508190509291505056fea26469706673582212202923bc7c5283d2701d87ad607da43cbf4fed90046cf8f001d6687c0fcfe0358864736f6c63430008110033",
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

// GetMessageHash is a free data retrieval call binding the contract method 0x082d8f29.
//
// Solidity: function getMessageHash(address receiver, string message, uint256 amount) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCaller) GetMessageHash(opts *bind.CallOpts, receiver common.Address, message string, amount *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "getMessageHash", receiver, message, amount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x082d8f29.
//
// Solidity: function getMessageHash(address receiver, string message, uint256 amount) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinSession) GetMessageHash(receiver common.Address, message string, amount *big.Int) ([32]byte, error) {
	return _Exchangecoin.Contract.GetMessageHash(&_Exchangecoin.CallOpts, receiver, message, amount)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x082d8f29.
//
// Solidity: function getMessageHash(address receiver, string message, uint256 amount) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCallerSession) GetMessageHash(receiver common.Address, message string, amount *big.Int) ([32]byte, error) {
	return _Exchangecoin.Contract.GetMessageHash(&_Exchangecoin.CallOpts, receiver, message, amount)
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

// WithdrawMoney is a paid mutator transaction binding the contract method 0x48bf759b.
//
// Solidity: function withdrawMoney(address receiver, string message, uint256 amount, bytes signature) returns()
func (_Exchangecoin *ExchangecoinTransactor) WithdrawMoney(opts *bind.TransactOpts, receiver common.Address, message string, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Exchangecoin.contract.Transact(opts, "withdrawMoney", receiver, message, amount, signature)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0x48bf759b.
//
// Solidity: function withdrawMoney(address receiver, string message, uint256 amount, bytes signature) returns()
func (_Exchangecoin *ExchangecoinSession) WithdrawMoney(receiver common.Address, message string, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Exchangecoin.Contract.WithdrawMoney(&_Exchangecoin.TransactOpts, receiver, message, amount, signature)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0x48bf759b.
//
// Solidity: function withdrawMoney(address receiver, string message, uint256 amount, bytes signature) returns()
func (_Exchangecoin *ExchangecoinTransactorSession) WithdrawMoney(receiver common.Address, message string, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Exchangecoin.Contract.WithdrawMoney(&_Exchangecoin.TransactOpts, receiver, message, amount, signature)
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
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventReceiverMoney is a free log retrieval operation binding the contract event 0xb0b6cfca2b3037ba69e8256ed0cf4e16b8d93ac32f69acb036c2c49576886ac0.
//
// Solidity: event eventReceiverMoney(address addr, uint256 amount)
func (_Exchangecoin *ExchangecoinFilterer) FilterEventReceiverMoney(opts *bind.FilterOpts) (*ExchangecoinEventReceiverMoneyIterator, error) {

	logs, sub, err := _Exchangecoin.contract.FilterLogs(opts, "eventReceiverMoney")
	if err != nil {
		return nil, err
	}
	return &ExchangecoinEventReceiverMoneyIterator{contract: _Exchangecoin.contract, event: "eventReceiverMoney", logs: logs, sub: sub}, nil
}

// WatchEventReceiverMoney is a free log subscription operation binding the contract event 0xb0b6cfca2b3037ba69e8256ed0cf4e16b8d93ac32f69acb036c2c49576886ac0.
//
// Solidity: event eventReceiverMoney(address addr, uint256 amount)
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

// ParseEventReceiverMoney is a log parse operation binding the contract event 0xb0b6cfca2b3037ba69e8256ed0cf4e16b8d93ac32f69acb036c2c49576886ac0.
//
// Solidity: event eventReceiverMoney(address addr, uint256 amount)
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
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEventWithDrawMoney is a free log retrieval operation binding the contract event 0xcf3bb011b58dadd278f376bd3eca311ad5135ee846d6efab8eff90dd0e1ddf0d.
//
// Solidity: event eventWithDrawMoney(address addr, bytes signature)
func (_Exchangecoin *ExchangecoinFilterer) FilterEventWithDrawMoney(opts *bind.FilterOpts) (*ExchangecoinEventWithDrawMoneyIterator, error) {

	logs, sub, err := _Exchangecoin.contract.FilterLogs(opts, "eventWithDrawMoney")
	if err != nil {
		return nil, err
	}
	return &ExchangecoinEventWithDrawMoneyIterator{contract: _Exchangecoin.contract, event: "eventWithDrawMoney", logs: logs, sub: sub}, nil
}

// WatchEventWithDrawMoney is a free log subscription operation binding the contract event 0xcf3bb011b58dadd278f376bd3eca311ad5135ee846d6efab8eff90dd0e1ddf0d.
//
// Solidity: event eventWithDrawMoney(address addr, bytes signature)
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

// ParseEventWithDrawMoney is a log parse operation binding the contract event 0xcf3bb011b58dadd278f376bd3eca311ad5135ee846d6efab8eff90dd0e1ddf0d.
//
// Solidity: event eventWithDrawMoney(address addr, bytes signature)
func (_Exchangecoin *ExchangecoinFilterer) ParseEventWithDrawMoney(log types.Log) (*ExchangecoinEventWithDrawMoney, error) {
	event := new(ExchangecoinEventWithDrawMoney)
	if err := _Exchangecoin.contract.UnpackLog(event, "eventWithDrawMoney", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
