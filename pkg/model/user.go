package model

import (
	"fmt"

	"github.com/arashi87/gin-template/pkg/common"
	"github.com/gin-gonic/gin"
)

type UserModel struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"column:name;not null" json:"name" binding:"required,max=15"`
	Email    string `gorm:"column:email;not null" json:"email" binding:"required,email"`
	Password string `gorm:"column:password;not null" json:"password"`
}

func (user *UserModel) Create(ctx *gin.Context) error {
	if result := common.DB.Create(user); result.Error != nil {
		return result.Error
	}
	fmt.Println("Create user")

	return nil
}

func (user *UserModel) Validate(ctx *gin.Context) error {
	if err := ctx.BindJSON(user); err != nil {
		return err
	}

	return nil
}
