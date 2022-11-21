package contract

import (

	"fmt"
	"log"
	"math/big"
	"servercoin/exchangecoin"
	"servercoin/utils"
    "servercoin/dto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
    "servercoin/service"
    "github.com/ethereum/go-ethereum/common/hexutil"
)

var (

	Mycontract *exchangecoin.Exchangecoin
    MycontractAddress string
)

func CreateNewContract(client *ethclient.Client) {

	contractAddress := utils.ReadContract(1)
    MycontractAddress = contractAddress
	address := common.HexToAddress(contractAddress)
    instance, err := exchangecoin.NewExchangecoin(address, client)
    if err != nil {
        
        fmt.Println()
        log.Println("Error create contract. Please check your contract's address!!!")
        fmt.Println()
        return
    }
	Mycontract = instance
    fmt.Println()
    fmt.Printf("Create contract at address %v successfully!!!\n", contractAddress)
    fmt.Println()
}

func Deploy(auth *bind.TransactOpts, client *ethclient.Client) {

    address, tx, instance, err := exchangecoin.DeployExchangecoin(auth, client)
    if err != nil {

        fmt.Println()
        log.Println("Error deploy contract. Please check your contract's source code!!!")
        fmt.Println()
        return
    }
    Mycontract = instance
    fmt.Println()
    fmt.Println("Deploy successfully. Your address of smart contract : ", address.Hex())
    fmt.Println("Your transaction hash : ", tx.Hash().Hex())
    fmt.Printf("Gas Price of Transaction : %v wei\n", tx.GasPrice())
    fmt.Println()
    MycontractAddress = address.Hex()
}

func WithdrawMoney(auth *bind.TransactOpts, receiver common.Address, message string, amount *big.Int, network *big.Int, signature []byte) {

    if Mycontract == nil {

        fmt.Println()
        fmt.Println("You have not created a contract!!!")
        fmt.Println()
        return
    }
    tx, err := Mycontract.WithdrawMoney(auth, receiver, message, amount, network, signature)
    if err != nil {

        fmt.Println()
        log.Printf("Error Withdraw Money : %v. Please check your signature and data!!!\n", err)
        fmt.Println()
        return
    }

    fmt.Println()
    value := new(big.Float).Quo(new(big.Float).SetInt(amount), new(big.Float).SetFloat64(1e18))
    fmt.Printf("You withdraw money successfully. You withdraw %v ether from contract %v !!!\n", value, MycontractAddress)
    fmt.Println("Your Transaction Hash : ", tx.Hash().Hex())
    fmt.Printf("Gas Price of Transaction : %v wei\n", tx.GasPrice())
    fmt.Println()
}

func GetBalance() {

    if Mycontract == nil {

        fmt.Println()
        fmt.Println("You have not created a contract!!!")
        fmt.Println()
        return
    }
    balance, err := Mycontract.GetBalance(&bind.CallOpts{})
    if err != nil {

        fmt.Println()
        log.Printf("Error getBalance of contract : %v. Please check your contract's address!!!\n", err)
        fmt.Println()
        return;
    }

    fmt.Println()
    fmt.Printf("Balance of Contract : %v wei\n", balance)
    balanceEther := new(big.Float)
    fmt.Printf("Balance of Contract : %v ether\n", balanceEther.Quo(new(big.Float).SetInt(balance), new(big.Float).SetFloat64(1e18)))
    fmt.Println()
}

func ReceiveBalance(auth *bind.TransactOpts) {

    if Mycontract == nil {

        fmt.Println()
        fmt.Println("You have not created a contract!!!")
        fmt.Println()
        return
    }
    tx, err := Mycontract.ReceiveMoney(auth)
    if err != nil {

        fmt.Println()
        log.Printf("Error send balance to contract : %v. Please check your contract's address!!!\n", err)
        fmt.Println()
        return
    }

    fmt.Println()
    value := new(big.Float).Quo(new(big.Float).SetInt(tx.Value()), new(big.Float).SetFloat64(1e18))
    fmt.Printf("You sent %v ether to contract %v!!!\n", value, MycontractAddress)
    fmt.Println("Transaction Hash :", tx.Hash().Hex())
    fmt.Printf("Transaction Value : %v ether\n", value)
    fmt.Println("Transaction Nonce : ", tx.Nonce())
    fmt.Printf("Gas Price of Transaction : %v wei\n", tx.GasPrice())
    fmt.Println()
}

func GetMessageHash(receiver common.Address, message string, amount *big.Int, network *big.Int) {

    if Mycontract == nil {

        fmt.Println()
        fmt.Println("You have not created a contract!!!")
        fmt.Println()
        return
    }
    mess, err := Mycontract.GetMessageHash(&bind.CallOpts{}, receiver, message, amount, network)
    if err != nil {

        fmt.Println()
        log.Printf("Error get MessageHash : %v. Please check your contract's address!!!\n", err)
        fmt.Println()
        return
    }
    messHash := mess[:]
    fmt.Println()
    fmt.Println("Message Hash : ", common.BytesToHash(messHash).Hex())
    fmt.Println()
}

func CheckMyContract() {

    for {

        if Mycontract != nil {

            return;
        }
    }
}

func ReadEventReceiverMoney(client *ethclient.Client) {

    CheckMyContract()
    logs := make(chan *exchangecoin.ExchangecoinEventReceiverMoney)
    sub, err := Mycontract.WatchEventReceiverMoney(&bind.WatchOpts{}, logs)
    if err != nil {

        fmt.Println()
        log.Println("Error subcribe event receive money : ", err)
        fmt.Println()
        return
    }

    fmt.Println()
    fmt.Printf("Listening Event Receive Money at contract %v .....\n", MycontractAddress)
    fmt.Println()
       
    for {
    
        select {
            
            case err := <-sub.Err():
                fmt.Println()
                log.Println("Event Receive Money : ", err)
                fmt.Println()
            case vLog := <-logs:
                fmt.Println()
                fmt.Println("This is new Event Receive Money!!!")
                balanceEther := new(big.Float).Quo(new(big.Float).SetInt(vLog.Amount), new(big.Float).SetFloat64(1e18))
                fmt.Printf("Address %v sent %v ether to contract!!!\n", vLog.Addr.Hex(), balanceEther)
                fmt.Println()

                message := utils.RandomMessage()
                data := &dto.Exchange{
            
                    Address: vLog.Addr.Hex(),
                    Amount: vLog.Amount.String(),
                    Message: message,
                    Signature: GenerateSignature(vLog.Addr, message, vLog.Amount.String()),
                    Withdrawn: false,
                }
                service.Serv.SaveEventReceiveMoney(data)
        }
    }
}

func ReadEventWithDrawMoney(client *ethclient.Client) {

    CheckMyContract()
    logs := make(chan *exchangecoin.ExchangecoinEventWithDrawMoney)
    sub, err := Mycontract.WatchEventWithDrawMoney(&bind.WatchOpts{}, logs)
    if err != nil {

        fmt.Println()
        log.Println("Error subcribe event withDraw money : ", err)
        return
    }

    fmt.Println()
    fmt.Printf("Listening Event Withdraw Money at contract %v .....\n", MycontractAddress)
    fmt.Println()

    for {

        select {

            case err := <-sub.Err():
                fmt.Println()
                log.Println("Event With Draw Money : ", err)
                fmt.Println()
            case vlog := <-logs:
                fmt.Println()
                fmt.Println("This is new Event WithDraw Money!!!")
                fmt.Printf("Address %v withdraw money!!!\n", vlog.Addr.Hex())
                fmt.Println()

                data := &dto.Exchange{

                    Address: vlog.Addr.Hex(),
                    Message: "",
                    Signature: hexutil.Encode(vlog.Signature),
                    Withdrawn: true,
                }
                service.Serv.SaveEventWithdrawMoney(data)
        }
    }
}