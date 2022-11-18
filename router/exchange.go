package router

import (

	"servercoin/handler"
	"servercoin/service"
	"github.com/gin-gonic/gin"
)

var (

	Handlers handler.ExchangeHandler
)

func ExchangeRouter(router *gin.Engine) {

	service.Init()
	Handlers = handler.NewExchangeHandler(service.Serv)
	router.GET("/signature", Handlers.GetExchange())
	router.POST("/signature", Handlers.CreateExchange())
	router.GET("/", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{

			"data" : "joker",
		})
	})
}