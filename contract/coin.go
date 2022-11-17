package contract

import (

	"bufio"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"servercoin/utils"
	"github.com/ethereum/go-ethereum/common"
)

var (

	in = bufio.NewReader(os.Stdin)
)

func handleChoice(choice int) {

	SetNonce()
	SetValue(big.NewInt(0))
	if choice == 1 {

		Deploy(Auth, Client)
	} else if choice == 2 {

		CreateNewContract(Client)
	} else if choice == 3 {

		GetBalance()
	} else if choice == 4 {

		var input float64
		fmt.Println("Enter number of ether you want to send to your contract : ")
		fmt.Fscan(in, &input)
		number := big.NewFloat(input)
		coin := new(big.Float)
		coin.SetInt(big.NewInt(1e18))
		bigval := number.Mul(number, coin)
		money := new(big.Int)
		bigval.Int(money)
		SetValue(money)
		ReceiveBalance(Auth)
	} else if choice == 5 {

		message := utils.RandomMessage()
		address := common.HexToAddress(utils.ReadPublicKey(2))
		amount := big.NewInt(rand.Int63n(1e18))
		signature := GenerateSignature(address, message, amount.String())
		fmt.Println("Address : ", address)
		fmt.Println("Message : ", message)
		fmt.Println("Amount : ", amount)
		fmt.Println("Signature : ", signature)
	} else if choice == 6 {

		ethHex := utils.ReadETHHash(1)
		ethHash := common.HexToHash(ethHex)
		signatureString := utils.ReadSignature(1)
		signature := common.FromHex(signatureString)
		publicKey := VerifySignature(ethHash, signature)
		fmt.Println("PublicKey : ", publicKey.Hex())
		fmt.Println(publicKey.Hex() == utils.ReadPublicKey(1))
	} else if choice == 7 {

		publicAddr := common.HexToAddress(utils.ReadPublicKey(2))
		amount := big.NewInt(1000000000000000000)
		message := "2smTyVsGd5Xav0yu99ZAMPTA7z7s575klKiz9pyKl17ltLSvQmntzYlkmifsd2X28"
		signature := common.FromHex(utils.ReadSignature(1))
		WithdrawMoney(Auth, publicAddr, message, amount, signature)
	} else if choice == 8 {

		message := "2smTyVsGd5Xav0yu99ZAMPTA7z7s575klKiz9pyKl17ltLSvQmntzYlkmifsd2X28"
		address := common.HexToAddress(utils.ReadPublicKey(2))
		// amount := rand.Int63n(1e18)
		GetMessageHash(address, message, big.NewInt(1000000000000000000))
		fmt.Println("Address : ", address)
		fmt.Println("Message : ", message)
		fmt.Println("Amount : ", 1000000000000000000)
	} else {

		fmt.Println("Input Invalid!!! Please enter again.")
	}
}

func Control() {

	fmt.Println("*******   Welcome to contract's control.   ********")
	for {

		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Println("1 : Deploy Contract.")
		fmt.Println("2 : Create New Contract.")
		fmt.Println("3 : GetBalance Contract.")
		fmt.Println("4 : Send ETH to Contract.")
		fmt.Println("5 : Generate Signature.")
		fmt.Println("6 : Verify Signature.")
		fmt.Println("7 : withDraw Money.")
		fmt.Println("8 : Get MessageHash.")
		fmt.Println()
		fmt.Fscan(in, &choice)
		handleChoice(choice)
		fmt.Println("----------------------------------------------------------------------------------")
	}
}
