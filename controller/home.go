package controller

import (
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context){
	c.JSON(200, gin.H{
		"status" : 200,
		"message": "API Homepage of Expense app",
	})
}