package controller

import (
	"github.com/gin-gonic/gin"
)

// HomePage Controller for the home page of API
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  200,
		"message": "API Homepage of Expense app",
	})
}
