package contract

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"servercoin/exchangecoin"
	"servercoin/utils"
)

var (

	Mycontract *exchangecoin.Exchangecoin
)

func CreateNewContract(auth *bind.TransactOpts, client *ethclient.Client) {

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