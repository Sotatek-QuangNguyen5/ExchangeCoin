package handler

import (

	"net/http"
	"servercoin/dto"
	"servercoin/errs"
	"servercoin/service"
	"github.com/gin-gonic/gin"
)


type ExchangeHandler struct {

	service service.ExchangeService
}

func NewExchangeHandler(service service.ExchangeService) ExchangeHandler {

	return ExchangeHandler{

		service: service,
	}
}

func (e ExchangeHandler) GetExchange() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var req = new(dto.Exchange)
		err := ctx.ShouldBindJSON(req)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		res, er := e.service.GetExchange(req.Address)
		if er != nil {

			WriteError(ctx, er)
			return
		}
		WriteRespon(ctx, http.StatusOK, res)
	}
}

func (e ExchangeHandler) CreateExchange() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var req = new(dto.Exchange)
		err := ctx.ShouldBindJSON(req)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		er := e.service.CreateExchange(req)
		if er != nil {

			WriteError(ctx, er)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.ResponsiveCreateSuccess("Signature"))
	}
}