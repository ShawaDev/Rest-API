package main

import (
	"github.com/ShawaDev/Rest-API/initializers"
	"github.com/ShawaDev/Rest-API/pkg/handler"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default()

	router.GET("/", handler.DaysUntill)

	router.Run()
}
