package router

import (
	"servercoin/config"
	"servercoin/middleware"

	"github.com/gin-gonic/gin"
)

func Start() {

	config.InitDB()
	router := gin.Default()
	router.Use(middleware.Cors())
	ExchangeRouter(router)
	router.Run(config.GetPort())
}