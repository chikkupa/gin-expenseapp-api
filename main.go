package main

import (
	"gin-expenseapp-api/controller"
	"gin-expenseapp-api/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	router.GET("/", controller.HomePage)
	router.POST("/login", controller.Login)

	authorized := router.Group("/")
	
	// Authenticated APIs
	authorized.Use(middleware.Authenticate())
	{
		transaction := authorized.Group("transaction");
		transaction.POST("/credit", controller.Credit)

	}

	router.NoRoute(controller.NotFound)

	router.Run(":8080")
}