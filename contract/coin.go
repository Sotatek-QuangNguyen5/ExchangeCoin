package contract

import (
	"fmt"

	"bufio"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

var (
	in = bufio.NewReader(os.Stdin)
)

func handleChoice(choice int) {

	SetNonce()
	SetValue(0)
	if choice == 1 {

		Deploy(Auth, Client)
	} else if choice == 2 {

		CreateNewContract(Auth, Client)
	} else if choice == 3 {

		// GetBalance()
	} else if choice == 4 {

		SetValue(1e18)
		// ReceiverBalance(Auth)
	} else if choice == 5 {

		message := "oke"
		GetMessageHash(message)
	} else if choice == 6 {

		GenerateSignature()
	} else if choice == 7 {

		signer := common.HexToAddress("0xC74cd3c249d2cEe4c11A894753213d2714962a7E")
		message := "oke"
		signature := common.FromHex("0xf3deaad8dd68d10c98051a10ba5309ec865d6323b79c56f4eca400247cd7f734060a40daad21363663c256eeed4f925c220f2edf135921e79526c798c75e77ad1c")
		// fmt.Println(signature)
		Verify(signer, message, signature)
	} else if choice == 8 {

		ethHash := common.HexToHash("0xb2767c9a44e1f883e8d99170ee032fcca1d248b6744609b78a9f0a11658329d4")
		signature := common.FromHex("0xf3deaad8dd68d10c98051a10ba5309ec865d6323b79c56f4eca400247cd7f734060a40daad21363663c256eeed4f925c220f2edf135921e79526c798c75e77ad1c")
		// fmt.Println(signature)
		Recover(ethHash, signature)
	} else {

		fmt.Println("Input Invalid!!! Please enter again.")
	}
}

func Coin() {

	Init()
	for {

		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Println("1 : Deploy Contract.")
		fmt.Println("2 : Create New Contract.")
		fmt.Println("3 : GetBalance Contract.")
		fmt.Println("4 : Send ETH to Contract.")
		fmt.Println("5 : Get Message Hash.")
		fmt.Println("6 : Generate Signature.")
		fmt.Println("7 : Verify Signature.")
		fmt.Println("8 : Recover Signature.")
		fmt.Fscan(in, &choice)
		handleChoice(choice)
		fmt.Println()
	}
}
