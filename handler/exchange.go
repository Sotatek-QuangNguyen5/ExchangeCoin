package handler

import (

	"net/http"
	"servercoin/contract"
	"servercoin/dto"
	"servercoin/errs"
	"servercoin/service"
	"servercoin/utils"

	"github.com/ethereum/go-ethereum/common"
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
		address := common.HexToAddress(req.Address)
		_, ethHash := contract.PacketWithEth(address.Bytes(), []byte(req.Message))
		clientAddress := contract.VerifySignature(ethHash, common.FromHex(req.Signature))
		// fmt.Println(req)
		// fmt.Println(clientAddress.Hex())
		if clientAddress.Hex() != req.Address {

			WriteRespon(ctx, http.StatusOK, gin.H{

				"Message" : "You are not address " + req.Address,
			})
			return;
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
		req.Message = utils.RandomMessage()
		addr := common.HexToAddress(req.Address)
		req.Signature = contract.GenerateSignature(addr, req.Message, req.Amount)
		er := e.service.CreateExchange(req)
		if er != nil {

			WriteError(ctx, er)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.ResponsiveCreateSuccess("Signature"))
	}
}