package controller

import (
	"kryptomind-web3-service/app/services"
	er "kryptomind-web3-service/app/utils/errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetBalance(ctx * gin.Context){
	addr :=  ctx.Param("address")
	if addr == " "{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": er.BadRequestError("Empty Address"),
		})
	}
	bal , err :=  services.BalanceDetails(addr)
	if err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{
			"message": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"Balance": bal,
	})

}