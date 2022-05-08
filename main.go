package main

import (
	"net/http"

	"github.com/arashi87/gin-template/pkg/setting"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	r.Run(setting.CONFIG.Address + ":" + setting.CONFIG.Port)
}
