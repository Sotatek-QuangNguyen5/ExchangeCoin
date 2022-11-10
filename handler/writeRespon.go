package handler

import (

	"servercoin/errs"
	"github.com/gin-gonic/gin"
)



func WriteError(ctx *gin.Context, err *errs.AppError) {

	ctx.AbortWithStatusJSON(err.Code, err)
}

func WriteRespon(ctx *gin.Context, status int, data interface{}) {

	ctx.JSON(status, data)
}