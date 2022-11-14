package main

import (
	"servercoin/contract"
	"servercoin/router"
)




func main() {

	contract.Init()
	go func ()  {
	
		contract.Coin()
	}()

	go func ()  {
		
		contract.ReadEventLog(contract.Client)
	}()
	router.Start()
}