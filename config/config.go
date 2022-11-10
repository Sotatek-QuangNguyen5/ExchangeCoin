package config

import (

	"log"
	"github.com/joho/godotenv"
)


func GetPort() string {

	myEvn, err := godotenv.Read("./.env")
	if err != nil {

		log.Fatal("Error read .env")
	}
	return myEvn["PORT"]
}