package controller

import (
	"net/http"

	"github.com/arashi87/gin-template/pkg/model"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var user model.UserModel

	// validate request data
	if err := user.Validate(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// create user
	if err := user.Create(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
