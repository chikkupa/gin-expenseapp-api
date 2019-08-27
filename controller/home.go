package controller

import (
	"gin-expenseapp-api/model"
	"github.com/gin-gonic/gin"
)

// HomePage Controller for the home page of API
func HomePage(c *gin.Context) {
	c.JSON(200, model.Response(model.Success, "Home page for Expense App API", map[string]string{
		"version" : "1.0.0",
	}))
}
