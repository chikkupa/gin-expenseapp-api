package main

import (
	"github.com/gin-gonic/gin"
	"gin-expenseapp/controller"
)

func main() {
	router := gin.Default()

	router.GET("/", controller.HomePage)
	router.POST("/credit", controller.Credit)
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")

		c.JSON(200, gin.H{
			"id" : id,
			"status" : 200,
			"message": "API Homepage of Expense app",
		})
	})
	router.Run(":8080")
}