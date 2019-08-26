package controller

import (
	"gin-expenseapp-api/model"
	"github.com/gin-gonic/gin"
)

// NotFound Controller for page not found
func NotFound(c *gin.Context) {
	c.JSON(404, model.Response(model.NotFound, "Page Not Found!", nil))
}