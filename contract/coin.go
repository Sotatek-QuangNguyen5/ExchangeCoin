package contract

import (

	"fmt"
	"servercoin/utils"
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

		v := GenerateSignature("1", "2", "3")
		_ = v
	} else if choice == 7 {

		address := utils.ReadPublicKey(1)
		signer := common.HexToAddress(address)
		message := "oke"
		utils.RandomMessage()
		signatureString := utils.ReadSignature(1)
		signature := common.FromHex(signatureString)
		// fmt.Println(signature)
		Verify(signer, message, signature)
	} else if choice == 8 {

		ethHash := common.HexToHash("0xb2767c9a44e1f883e8d99170ee032fcca1d248b6744609b78a9f0a11658329d4")
		signatureString := utils.ReadSignature(1)
		signature := common.FromHex(signatureString)
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
