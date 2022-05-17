package controller

import (
	"net/http"
	"strconv"

	"github.com/arashi87/gin-template/pkg/common"
	"github.com/arashi87/gin-template/pkg/model"
	"github.com/gin-gonic/gin"
)

// @Summary User create
// @Version 1.0
// @Description User create
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param data body model.UserModel true "request body"
// @Success 200 {string} json "{"msg":"ok"}"
// @Failure 400 {string} json "{"msg":"error reason"}"
// @Router /user [post]
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

// @Summary User retrive
// @Version 1.0
// @Description User retrive
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param uid path string true "uid"
// @Param Authorization header string false "Bearer token"
// @Success 200 {object} model.UserModel "{"msg":"ok"}"
// @Failure 400 {string} json "{"msg":"error reason"}"
// @Router /user/{uid} [get]
func RetrieveUser(ctx *gin.Context) {
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

// @Summary User retrive
// @Version 1.0
// @Description User retrive
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param uid path string true "uid"
// @Param Authorization header string false "Bearer token"
// @Success 200 {object} model.UserModel "{"msg":"ok"}"
// @Failure 400 {string} json "{"msg":"error reason"}"
// @Router /user/{uid} [put]
func UpdateUser(ctx *gin.Context) {
	var user model.UserModel

	// validate request data
	if err := user.Validate(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// update user
	if err := user.Update(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// @Summary User delete
// @Version 1.0
// @Description User delete
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param uid path string true "uid"
// @Param Authorization header string false "Bearer token"
// @Success 200 {object} model.UserModel "{"msg":"ok"}"
// @Failure 400 {string} json "{"msg":"error reason"}"
// @Router /user/{uid} [delete]
func DeleteUser(ctx *gin.Context) {
	// extract uid from param
	uid, err := strconv.Atoi(ctx.Param("uid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// delete user
	user := model.UserModel{ID: uid}
	if err := user.Delete(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
