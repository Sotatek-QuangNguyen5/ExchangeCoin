package contract

import (
	"fmt"
	"log"

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

func GetMessageHash(message string) {

    res, err := Mycontract.GetMessageHash(&bind.CallOpts{}, message)
    if err != nil {

        log.Fatal(err)
    }

    fmt.Print("Message Hash : {")
    for _, value := range res {

        fmt.Print(value, ", ")
    }
    fmt.Println("}")
}

func Verify(signer common.Address, message string, signature []byte) {

    res, err := Mycontract.Verify(&bind.CallOpts{}, signer, message, signature)
    if err != nil {

        log.Fatal("Error Verify : ", err)
    }
    fmt.Println("Result : ", res)
}

func Recover(ethHash [32]byte, signature []byte) {

    res, err := Mycontract.RecoverSigner(&bind.CallOpts{}, ethHash, signature)
    if err != nil {

        log.Fatal("Error Recover : ", err)
    }
    fmt.Println("Address : ", res.Hex())
}