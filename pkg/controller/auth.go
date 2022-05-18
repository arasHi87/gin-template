package controller

import (
	"net/http"

	"github.com/arashi87/gin-template/pkg/common"
	"github.com/arashi87/gin-template/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthInfo struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary User authorization
// @Version 1.0
// @Description User login
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param data body AuthInfo true "request body"
// @Success 200 {string} json "{"msg":"ok"}"
// @Failure 400 {string} json "{"msg":"error reason"}"
// @Failure 403 {string} json "{"msg":"user name or password wrong"}"
// @Router /auth [post]
func AuthLogin(ctx *gin.Context) {
	var authInfo AuthInfo
	var user model.UserModel

	// bind auth info
	if err := ctx.BindJSON(&authInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// check if user exist
	result := common.DB.Where(&model.UserModel{Name: authInfo.Name}).Find(&user)
	if result.Error != nil {
		common.Logger.WithFields(logrus.Fields{
			"type": "user auth error",
		}).Error(result.Error.Error())
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"msg": result.Error.Error()})
		return
	}

	// compare password
	if common.CheckPassword(authInfo.Password, user.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "user name or password not match"})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusForbidden, gin.H{"msg": "user name or password wrong"})
		return
	}

	// user exist, generate jwt token
	token, err := common.GenToken(user.Name, user.Email, user.ID)
	if err != nil {
		common.Logger.WithFields(logrus.Fields{
			"type": "user auth error",
		}).Error(result.Error.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": result.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "ok", "data": map[string]string{"token": token}})
}
