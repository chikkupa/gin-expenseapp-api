package main

import (
	"github.com/gin-gonic/gin"
	"gin-expenseapp/controller"
)

func main() {
	router := gin.Default()

	router.GET("/", controller.HomePage)
	router.POST("/credit", controller.Credit)
	
	router.Run(":8080")
}