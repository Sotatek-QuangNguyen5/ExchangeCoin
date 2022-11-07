package contract

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"servercoin/exchangecoin"
	"servercoin/readfile"
)

var (

	Mycontract *exchangecoin.Exchangecoin
)

func CreateNewContract(auth *bind.TransactOpts, client *ethclient.Client) {

	contractAddress := readfile.ReadContractOne()
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

func GetBalance() {

    balance, err := Mycontract.GetBalance(nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Balance of Contract : %v ether\n", balance.Div(balance, big.NewInt(1e18)))
}

func ReceiverBalance(auth *bind.TransactOpts) {

    balance, err := Mycontract.ReceiverBalance(auth)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Transaction Hash :", balance.Hash().Hex())
}

func Joker(auth *bind.TransactOpts) {

    Mycontract.
}