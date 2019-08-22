package main

import (
	"gin-expenseapp-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", controller.HomePage)
	router.POST("/credit", controller.Credit)

	router.Run(":8080")
}
