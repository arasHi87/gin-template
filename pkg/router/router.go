package router

import (
	_ "github.com/arashi87/gin-template/docs"
	"github.com/arashi87/gin-template/pkg/controller"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", controller.GetHealth)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}