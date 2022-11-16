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
    "github.com/ethereum/go-ethereum/accounts/abi"
	"servercoin/exchangecoin"
	"servercoin/utils"
    "strings"
    "github.com/mitchellh/mapstructure"
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

func WithdrawMoney(auth *bind.TransactOpts, receiver common.Address, message string, amount *big.Int, signature []byte) {

    res, err := Mycontract.WithdrawMoney(auth, receiver, message, amount, signature)
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
    balanceEther := new(big.Float)
    fmt.Printf("Balance of Contract : %v ether\n", balanceEther.Quo(new(big.Float).SetInt(balance), new(big.Float).SetFloat64(1e18)))
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

    exchangecoin.ExchangecoinEventReceiverMoney
    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {

        log.Fatal(err)
    }

    contractAbi, err := abi.JSON(strings.NewReader(string(exchangecoin.ExchangecoinABI)))
    if err != nil {
        log.Fatal(err)
    }

    eventData := struct {
        
        addr common.Address
        amount *big.Int
    }{}
       
    for {
    
        select {
            
            case err := <-sub.Err():
                log.Fatal(err)
            case vLog := <-logs:
                //fmt.Println("This is Event Log : ", vLog.Topics) // pointer to event log
                //Topics(vLog.Topics)
                event, err := contractAbi.Unpack("eventReceiverMoney", vLog.Data)
                if err != nil {
                    log.Fatal(err)
                }
                mapstructure.Decode(event, &eventData)
                fmt.Println(eventData.addr)
                fmt.Println(eventData.addr)
                fmt.Println(eventData.amount)
        }
    }
}

func Topics(topics []common.Hash) {

    for i, v := range topics {

        fmt.Println("Topic ", i, " : ", v.Hex())
    }
}

func GetMessageHash(receiver common.Address, message string, amount *big.Int) {

    mess, err := Mycontract.GetMessageHash(&bind.CallOpts{}, receiver, message, amount)
    if err != nil {

        log.Fatal("Error get MessageHash : ", err)
    }
    messHash := mess[:]
    fmt.Println("Message Hash : ", common.BytesToHash(messHash).Hex())
}