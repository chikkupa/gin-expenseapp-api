package main

import (
	"io/ioutil"
	"bytes"
	"net/http"

	"gin-expenseapp-api/controller"
	"gin-expenseapp-api/middleware"
	"gin-expenseapp-api/library"
	"gin-expenseapp-api/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	router.Use(ValidateRequest)

	router.GET("/", controller.HomePage)
	router.POST("/login", controller.Login)

	authorized := router.Group("/")
	
	// Authenticated APIs
	authorized.Use(middleware.Authenticate())
	{
		transaction := authorized.Group("transaction");
		transaction.POST("/credit", controller.Credit)
		transaction.POST("/debit", controller.Debit)
	}

	router.NoRoute(controller.NotFound)

	router.Run(":8080")
}

// ValidateRequest Validate API request body
func ValidateRequest(c *gin.Context){
	if c.Request.Method != http.MethodGet {
		b, _ := ioutil.ReadAll(c.Request.Body)

		if(!library.IsJSON(string(b))){
			c.AbortWithStatusJSON(400, model.Response(model.BadRequestError, "Invalid JSON Request", nil))
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	}

	c.Next();
}