package router

import (
	_ "github.com/arashi87/gin-template/docs"
	"github.com/arashi87/gin-template/pkg/controller"
	"github.com/arashi87/gin-template/pkg/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	/* General router */
	{
		router.GET("/health", controller.GetHealth)
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		router.POST("/auth", controller.AuthLogin)
		router.POST("/user", controller.CreateUser)
	}

	/* User router */
	user := router.Group("/user", middleware.AuthMiddleware())
	{
		user.GET("/:uid", controller.RetrieveUser)
	}

	return router
}
