package model

import (
	"errors"
	"strconv"

	"github.com/arashi87/gin-template/pkg/common"
	"github.com/gin-gonic/gin"
)

type UserModel struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"column:name;not null" json:"name" binding:"required,max=15"`
	Email    string `gorm:"column:email;not null" json:"email" binding:"required,email"`
	Password string `gorm:"column:password;not null" json:"password"`
}

func (model *UserModel) Create(ctx *gin.Context) error {
	var user UserModel

	// hash password
	password, err := common.HashPassword(model.Password)
	if err != nil {
		return err
	}
	model.Password = string(password)

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

	return nil
}

func (model *UserModel) Update(ctx *gin.Context) error {
	user := ctx.MustGet("user").(UserModel)

	// check if user is self
	if uid, err := strconv.Atoi(ctx.Param("uid")); uid != user.ID || err != nil {
		return errors.New("permission denied")
	}

	// check is password need hash
	if model.Password != "" {
		password, err := common.HashPassword(model.Password)
		if err != nil {
			return err
		}
		model.Password = string(password)
	}

	record := common.DB.Model(&user).Updates(model)
	if err := record.Error; err != nil {
		return err
	}

	return nil
}

func (model *UserModel) Delete(ctx *gin.Context) error {
	user := ctx.MustGet("user").(UserModel)

	// check if user is self
	if uid, err := strconv.Atoi(ctx.Param("uid")); uid != user.ID || err != nil {
		return errors.New("permission denied")
	}

	// delete user
	record := common.DB.Delete(&UserModel{}, ctx.Param("uid"))
	if err := record.Error; err != nil {
		return err
	}

	return nil
}

func (model *UserModel) Validate(ctx *gin.Context) error {
	if err := ctx.BindJSON(model); err != nil {
		return err
	}

	return nil
}
