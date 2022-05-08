package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Health check
// @Version 1.0
// @Description Get service health
// @Tags health
// @Produce text/plain
// @Success 200 {string} json "{"msg":"ok"}"
// @Router /health [get]
func GetHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
