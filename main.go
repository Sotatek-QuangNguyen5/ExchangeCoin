package main

import (

	"fmt"
	"servercoin/contract"
	"servercoin/router"
	"time"
)

func main() {

	contract.Init()
	go func ()  {
	
		time.Sleep(1 * time.Second)
		fmt.Println()
		fmt.Println()
		contract.Control()
	}()
	go func() {

        contract.ReadEventReceiverMoney(contract.Client)
    }()
	go func()  {
	
		contract.ReadEventWithDrawMoney(contract.Client)
	}()
	router.Start()
}