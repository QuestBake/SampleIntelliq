package main

import (
	"github.com/gin-gonic/gin"

	"SampleIntelliq/app/approuter"
	"SampleIntelliq/app/config"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	approuter.AddRouters(router)
	config.Connect()
	router.Run()
}
