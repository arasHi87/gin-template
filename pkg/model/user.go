package model

import (
	"errors"
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

func (model *UserModel) Create(ctx *gin.Context) error {
	var user UserModel

	// check if user info has already exist
	record := common.DB.Where("name = ? or email = ?", model.Name, model.Email).Limit(1).Find(&user)
	if err := record.Error; err != nil {
		return err
	}

	if record.RowsAffected > 0 {
		return errors.New("user has already exist")
	}

	// create user
	if result := common.DB.Create(model); result.Error != nil {
		return result.Error
	}
	fmt.Println("Create user")

	return nil
}

func (model *UserModel) Validate(ctx *gin.Context) error {
	if err := ctx.BindJSON(model); err != nil {
		return err
	}

	return nil
}
