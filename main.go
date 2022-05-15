package main

import (
	"fmt"

	"github.com/arashi87/gin-template/pkg/common"
	"github.com/arashi87/gin-template/pkg/model"
	"github.com/arashi87/gin-template/pkg/router"
	"github.com/arashi87/gin-template/pkg/setting"
)

func DBAutoMigrate() {
	common.DB.AutoMigrate(&model.UserModel{})
	fmt.Println("Auto migrate all models")
}

// @title gin template
// @version 1.0
// @description a simple RESTful API service implement by gin
// @termsOfService http://swagger.io/terms/

// @contact.name arashi87
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// schemes http
func main() {
	// init database connection
	common.InitDatabase()

	// migrate models
	DBAutoMigrate()

	router := router.InitRouter()
	router.Run(setting.CONFIG.Address + ":" + setting.CONFIG.Port)
}
