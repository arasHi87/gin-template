package common

import (
	"github.com/gin-gonic/gin"
)

type Serializer interface {
	Create(ctx *gin.Context) error
	Retrive(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Validate(ctx *gin.Context) error
}
