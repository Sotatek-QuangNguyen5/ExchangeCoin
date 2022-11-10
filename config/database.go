package config

import (
	"fmt"
	"log"
	"servercoin/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (

	DB *gorm.DB
)

func connect() *gorm.DB {

	var myEnv map[string]string
	myEnv, err := godotenv.Read("./.env")

	if err != nil {

		log.Fatal("Error read env!!!")
	}

	var (

		db_name = myEnv["DB_NAME"]
		db_user = myEnv["DB_USER"]
		db_host = myEnv["DB_HOST"]
		db_password = myEnv["DB_PASSWORD"]
		db_port = myEnv["DB_PORT"]
	)
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", db_user, db_password, db_host, db_port, db_name) 
	// "user:pass@tcp(127.0.0.1:3306)/dbname
	// ?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		log.Fatal("Error database!!!", err)
	}
	db.AutoMigrate(&model.Exchange{})

	return db
}

func InitDB() {

	DB = connect()
}