package router

import (
	"servercoin/config"
	"servercoin/handler"
	"servercoin/repository"
	"servercoin/service"

	"github.com/gin-gonic/gin"
)

func ExchangeRouter(router *gin.Engine) {

	h := handler.NewExchangeHandler(service.NewExchangeService(repository.NewExchangeRepository(config.DB)))
	router.GET("/signature", h.GetExchange())
	router.POST("/signature", h.CreateExchange())
}