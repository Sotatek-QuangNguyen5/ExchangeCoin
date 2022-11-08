package main

import (
	"fmt"

	"bufio"
	"os"
	"servercoin/contract"
	"servercoin/initstate"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

var (

	in = bufio.NewReader(os.Stdin)
)

func handleChoice(choice int) {

	initstate.SetNonce()
	initstate.SetValue(0)
	if choice == 1 {

		contract.Deploy(initstate.Auth, initstate.Client)
	} else if choice == 2 {

		contract.CreateNewContract(initstate.Auth, initstate.Client)
	} else if choice == 3 {

		contract.GetBalance()
	} else if choice == 4 {

		initstate.SetValue(1e18)
		contract.ReceiverBalance(initstate.Auth)
	} else if choice == 5 {

		to := common.HexToAddress("0x9E4C79E571c85dF7C9AE41260c7F634aBD1f442F")
		amount := *big.NewInt(1e18)
		message := "oke"
		fmt.Println("Nonce : ", initstate.Auth.Nonce)
		contract.GetMessageHash(to, &amount, message, initstate.Auth.Nonce)
	} else if choice == 6 {

		initstate.GenerateSignature();
	} else {

		fmt.Println("Input Invalid!!! Please enter again.")
	}
}

func main() {

	initstate.Init()
	for {

		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Println("1 : Deploy Contract.")
		fmt.Println("2 : Create New Contract.")
		fmt.Println("3 : GetBalance Contract.")
		fmt.Println("4 : Send ETH to Contract.")
		fmt.Println("5 : Get Message Hash.")
		fmt.Println("6 : Generate Signature.")
		fmt.Fscan(in, &choice)
		handleChoice(choice)
		fmt.Println()
	}
}