package main

import (
	"pracSpace/restHandler_Gin/app/approuter"
	"pracSpace/restHandler_Gin/app/config"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	approuter.AddRouters(router)
	config.Connect()
	router.Run()
}
