package utils

import (

	"servercoin/config"
	"servercoin/repository"
	"servercoin/service"
)


func RandomMessage() {

	s := service.NewExchangeService(repository.NewExchangeRepository(config.DB))

	for i := 0; i < 32; i++ {

		
	}
}