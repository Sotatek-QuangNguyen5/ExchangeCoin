package contract

import (
	"fmt"
	"log"
	"math/big"
    "context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/core/types"
	"servercoin/exchangecoin"
	"servercoin/utils"
)

var (

	Mycontract *exchangecoin.Exchangecoin
)

func CreateNewContract(client *ethclient.Client) {

	contractAddress := utils.ReadContract(1)
	address := common.HexToAddress(contractAddress)
    instance, err := exchangecoin.NewExchangecoin(address, client)
    if err != nil {
        log.Fatal(err)
    }
	Mycontract = instance
}

func Deploy(auth *bind.TransactOpts, client *ethclient.Client) {

    address, tx, instance, err := exchangecoin.DeployExchangecoin(auth, client)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(address.Hex())
    fmt.Println(tx.Hash().Hex())
	Mycontract = instance
}

func WithdrawMoney(auth *bind.TransactOpts, receiver common.Address, message string, amount *big.Int, nonce *big.Int, signature []byte) {

    res, err := Mycontract.WithdrawMoney(auth, receiver, message, amount, nonce, signature)
    if err != nil {

        log.Fatal("Error Verify : ", err)
    }
    fmt.Println("Result : ", res.Hash().Hex())
}

func GetBalance() {

    balance, err := Mycontract.GetBalance(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Balance of Contract : %v wei\n", balance)
    fmt.Printf("Balance of Contract : %v ether\n", balance.Div(balance, big.NewInt(1e18)))
}


func ReceiveBalance(auth *bind.TransactOpts) {

    tx, err := Mycontract.ReceiveMoney(auth)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Transaction Hash :", tx.Hash().Hex())
}

func ReadEventLog(client *ethclient.Client) {

    addr := utils.ReadContract(1)
    contractAddress := common.HexToAddress(addr)
    query := ethereum.FilterQuery {

        Addresses: []common.Address{contractAddress},
    }

    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {

        log.Fatal(err)
    }

    fmt.Println("Joke Coin : ", contractAddress.Hex())
       
    for {
    
        select {
            
            case err := <-sub.Err():
                log.Fatal(err)
            case vLog := <-logs:
                fmt.Println("This is Event Log : ", vLog.Address.Hex()) // pointer to event log
        }
    }
}

func SetNumber(auth *bind.TransactOpts, x int64) {

    tx, err := Mycontract.SetNumber(auth, big.NewInt(x))
    if err != nil {

        log.Fatal("Error : ", err)
    }
    fmt.Println("Transaction : ", tx.Hash().Hex())
}

func GetNumber() {

    tx, err := Mycontract.GetNumber(&bind.CallOpts{})
    if err != nil {

        log.Fatal("Error : ", err)
    }
    fmt.Println("Transaction : ", tx)
}