package routes

import (
	"kryptomind-web3-service/app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Url_maps() {
	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Is Alive....!!")
	})
	v1 := r.Group("/api/v1/")
	{
		v1.GET("/getbalance/:address", controller.GetBalance)
	}
}
