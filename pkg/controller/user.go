package controller

import (
	"net/http"
	"strconv"

	"github.com/arashi87/gin-template/pkg/common"
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

func RetriveUser(ctx *gin.Context) {
	var user model.UserModel
	record := common.DB.Select("id", "name", "email").Where("id = ?", ctx.Param("uid")).Limit(1).Find(&user)

	if err := record.Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// return 404 if user not found
	if record.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "ok",
		"data": map[string]string{"id": strconv.Itoa(int(user.ID)), "name": user.Name, "email": user.Email}})
}
