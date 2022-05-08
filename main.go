package main

import (
	"github.com/arashi87/gin-template/pkg/router"
	"github.com/arashi87/gin-template/pkg/setting"
)

func main() {
	router := router.InitRouter()
	router.Run(setting.CONFIG.Address + ":" + setting.CONFIG.Port)
}
