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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"onSetNumber\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"getEthSignedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveMoney\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ethSignedMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdrawMoney\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611257806100606000396000f3fe6080604052600436106100a75760003560e01c80636d26ec18116100645780636d26ec18146101d057806397aba7f9146101da578063a5810c6e14610217578063a7bb580314610254578063f2c9ecd814610293578063fa540801146102be576100a7565b806302d05d3f146100ac57806312065fe0146100d757806323fd0e40146101025780633fb5c1cb1461012d57806358b77985146101565780636c9ee8d114610193575b600080fd5b3480156100b857600080fd5b506100c16102fb565b6040516100ce91906107f2565b60405180910390f35b3480156100e357600080fd5b506100ec61031f565b6040516100f99190610826565b60405180910390f35b34801561010e57600080fd5b50610117610327565b6040516101249190610826565b60405180910390f35b34801561013957600080fd5b50610154600480360381019061014f9190610881565b61032d565b005b34801561016257600080fd5b5061017d60048036038101906101789190610ac1565b61036e565b60405161018a9190610b8f565b60405180910390f35b34801561019f57600080fd5b506101ba60048036038101906101b59190610ac1565b6103f1565b6040516101c79190610b8f565b60405180910390f35b6101d86105d7565b005b3480156101e657600080fd5b5061020160048036038101906101fc9190610be0565b610667565b60405161020e91906107f2565b60405180910390f35b34801561022357600080fd5b5061023e60048036038101906102399190610c3c565b6106d6565b60405161024b9190610cce565b60405180910390f35b34801561026057600080fd5b5061027b60048036038101906102769190610ce9565b61070f565b60405161028a93929190610d4e565b60405180910390f35b34801561029f57600080fd5b506102a8610777565b6040516102b59190610826565b60405180910390f35b3480156102ca57600080fd5b506102e560048036038101906102e09190610d85565b610781565b6040516102f29190610cce565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600047905090565b60015481565b7f1f6641444d11c93a5ca4da333655135a1b21c94ed0a6ba2f156a0dfc4626a1683360405161035c91906107f2565b60405180910390a18060018190555050565b60008061037d878787876106d6565b9050600061038a82610781565b905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166103cd8286610667565b73ffffffffffffffffffffffffffffffffffffffff16149250505095945050505050565b600080610400878787876106d6565b9050600061040d82610781565b905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166104508286610667565b73ffffffffffffffffffffffffffffffffffffffff16146104a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049d90610e0f565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614610514576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050b90610e7b565b60405180910390fd5b47861115610557576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054e90610ee7565b60405180910390fd5b60003373ffffffffffffffffffffffffffffffffffffffff16876113889060405161058190610f38565b600060405180830381858888f193505050503d80600081146105bf576040519150601f19603f3d011682016040523d82523d6000602084013e6105c4565b606091505b5050905080935050505095945050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610665576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065c90610f99565b60405180910390fd5b565b6000806000806106768561070f565b9250925092506001868285856040516000815260200160405260405161069f9493929190610fb9565b6020604051602081039080840390855afa1580156106c1573d6000803e3d6000fd5b50505060206040510351935050505092915050565b6000848484846040516020016106ef94939291906110d8565b604051602081830303815290604052805190602001209050949350505050565b60008060006041845114610758576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074f9061116e565b60405180910390fd5b6020840151925060408401519150606084015160001a90509193909250565b6000600154905090565b60008160405160200161079491906111fb565b604051602081830303815290604052805190602001209050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006107dc826107b1565b9050919050565b6107ec816107d1565b82525050565b600060208201905061080760008301846107e3565b92915050565b6000819050919050565b6108208161080d565b82525050565b600060208201905061083b6000830184610817565b92915050565b6000604051905090565b600080fd5b600080fd5b61085e8161080d565b811461086957600080fd5b50565b60008135905061087b81610855565b92915050565b6000602082840312156108975761089661084b565b5b60006108a58482850161086c565b91505092915050565b6108b7816107d1565b81146108c257600080fd5b50565b6000813590506108d4816108ae565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61092d826108e4565b810181811067ffffffffffffffff8211171561094c5761094b6108f5565b5b80604052505050565b600061095f610841565b905061096b8282610924565b919050565b600067ffffffffffffffff82111561098b5761098a6108f5565b5b610994826108e4565b9050602081019050919050565b82818337600083830152505050565b60006109c36109be84610970565b610955565b9050828152602081018484840111156109df576109de6108df565b5b6109ea8482856109a1565b509392505050565b600082601f830112610a0757610a066108da565b5b8135610a178482602086016109b0565b91505092915050565b600067ffffffffffffffff821115610a3b57610a3a6108f5565b5b610a44826108e4565b9050602081019050919050565b6000610a64610a5f84610a20565b610955565b905082815260208101848484011115610a8057610a7f6108df565b5b610a8b8482856109a1565b509392505050565b600082601f830112610aa857610aa76108da565b5b8135610ab8848260208601610a51565b91505092915050565b600080600080600060a08688031215610add57610adc61084b565b5b6000610aeb888289016108c5565b955050602086013567ffffffffffffffff811115610b0c57610b0b610850565b5b610b18888289016109f2565b9450506040610b298882890161086c565b9350506060610b3a8882890161086c565b925050608086013567ffffffffffffffff811115610b5b57610b5a610850565b5b610b6788828901610a93565b9150509295509295909350565b60008115159050919050565b610b8981610b74565b82525050565b6000602082019050610ba46000830184610b80565b92915050565b6000819050919050565b610bbd81610baa565b8114610bc857600080fd5b50565b600081359050610bda81610bb4565b92915050565b60008060408385031215610bf757610bf661084b565b5b6000610c0585828601610bcb565b925050602083013567ffffffffffffffff811115610c2657610c25610850565b5b610c3285828601610a93565b9150509250929050565b60008060008060808587031215610c5657610c5561084b565b5b6000610c64878288016108c5565b945050602085013567ffffffffffffffff811115610c8557610c84610850565b5b610c91878288016109f2565b9350506040610ca28782880161086c565b9250506060610cb38782880161086c565b91505092959194509250565b610cc881610baa565b82525050565b6000602082019050610ce36000830184610cbf565b92915050565b600060208284031215610cff57610cfe61084b565b5b600082013567ffffffffffffffff811115610d1d57610d1c610850565b5b610d2984828501610a93565b91505092915050565b600060ff82169050919050565b610d4881610d32565b82525050565b6000606082019050610d636000830186610cbf565b610d706020830185610cbf565b610d7d6040830184610d3f565b949350505050565b600060208284031215610d9b57610d9a61084b565b5b6000610da984828501610bcb565b91505092915050565b600082825260208201905092915050565b7f696e76616c6964207369676e6174757265212121000000000000000000000000600082015250565b6000610df9601483610db2565b9150610e0482610dc3565b602082019050919050565b60006020820190508181036000830152610e2881610dec565b9050919050565b7f77726f6e67207265636569766572212121000000000000000000000000000000600082015250565b6000610e65601183610db2565b9150610e7082610e2f565b602082019050919050565b60006020820190508181036000830152610e9481610e58565b9050919050565b7f6e6f7420656e6f756768206d6f6e657921212100000000000000000000000000600082015250565b6000610ed1601383610db2565b9150610edc82610e9b565b602082019050919050565b60006020820190508181036000830152610f0081610ec4565b9050919050565b600081905092915050565b50565b6000610f22600083610f07565b9150610f2d82610f12565b600082019050919050565b6000610f4382610f15565b9150819050919050565b7f4e6f742061757468656e7469636174696f6e2121210000000000000000000000600082015250565b6000610f83601583610db2565b9150610f8e82610f4d565b602082019050919050565b60006020820190508181036000830152610fb281610f76565b9050919050565b6000608082019050610fce6000830187610cbf565b610fdb6020830186610d3f565b610fe86040830185610cbf565b610ff56060830184610cbf565b95945050505050565b60008160601b9050919050565b600061101682610ffe565b9050919050565b60006110288261100b565b9050919050565b61104061103b826107d1565b61101d565b82525050565b600081519050919050565b600081905092915050565b60005b8381101561107a57808201518184015260208101905061105f565b60008484015250505050565b600061109182611046565b61109b8185611051565b93506110ab81856020860161105c565b80840191505092915050565b6000819050919050565b6110d26110cd8261080d565b6110b7565b82525050565b60006110e4828761102f565b6014820191506110f48286611086565b915061110082856110c1565b60208201915061111082846110c1565b60208201915081905095945050505050565b7f696e76616c6964207369676e6174757265206c656e6774682121210000000000600082015250565b6000611158601b83610db2565b915061116382611122565b602082019050919050565b600060208201905081810360008301526111878161114b565b9050919050565b7f4261746d616e207673204a6f6b65720000000000000000000000000000000000600082015250565b60006111c4600f83611051565b91506111cf8261118e565b600f82019050919050565b6000819050919050565b6111f56111f082610baa565b6111da565b82525050565b6000611206826111b7565b915061121282846111e4565b6020820191508190509291505056fea26469706673582212208b3a597b8a05556a22a4eeedaf88b3d1deda09ebb4dac6102a0fb5ea8fc18f5864736f6c63430008110033",
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
// Solidity: function getMessageHash(address receiver, string message, uint256 amount, uint256 nonce) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCaller) GetMessageHash(opts *bind.CallOpts, receiver common.Address, message string, amount *big.Int, nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "getMessageHash", receiver, message, amount, nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0xa5810c6e.
//
// Solidity: function getMessageHash(address receiver, string message, uint256 amount, uint256 nonce) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinSession) GetMessageHash(receiver common.Address, message string, amount *big.Int, nonce *big.Int) ([32]byte, error) {
	return _Exchangecoin.Contract.GetMessageHash(&_Exchangecoin.CallOpts, receiver, message, amount, nonce)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xa5810c6e.
//
// Solidity: function getMessageHash(address receiver, string message, uint256 amount, uint256 nonce) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCallerSession) GetMessageHash(receiver common.Address, message string, amount *big.Int, nonce *big.Int) ([32]byte, error) {
	return _Exchangecoin.Contract.GetMessageHash(&_Exchangecoin.CallOpts, receiver, message, amount, nonce)
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

// Verify is a free data retrieval call binding the contract method 0x58b77985.
//
// Solidity: function verify(address receiver, string message, uint256 amount, uint256 nonce, bytes signature) view returns(bool)
func (_Exchangecoin *ExchangecoinCaller) Verify(opts *bind.CallOpts, receiver common.Address, message string, amount *big.Int, nonce *big.Int, signature []byte) (bool, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "verify", receiver, message, amount, nonce, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x58b77985.
//
// Solidity: function verify(address receiver, string message, uint256 amount, uint256 nonce, bytes signature) view returns(bool)
func (_Exchangecoin *ExchangecoinSession) Verify(receiver common.Address, message string, amount *big.Int, nonce *big.Int, signature []byte) (bool, error) {
	return _Exchangecoin.Contract.Verify(&_Exchangecoin.CallOpts, receiver, message, amount, nonce, signature)
}

// Verify is a free data retrieval call binding the contract method 0x58b77985.
//
// Solidity: function verify(address receiver, string message, uint256 amount, uint256 nonce, bytes signature) view returns(bool)
func (_Exchangecoin *ExchangecoinCallerSession) Verify(receiver common.Address, message string, amount *big.Int, nonce *big.Int, signature []byte) (bool, error) {
	return _Exchangecoin.Contract.Verify(&_Exchangecoin.CallOpts, receiver, message, amount, nonce, signature)
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
