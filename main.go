package main

import (
	"servercoin/contract"
	"servercoin/router"
)




func main() {

	go func ()  {
		
		contract.Coin()
	}()
	router.Start()
}