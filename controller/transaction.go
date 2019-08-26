package controller

import (
	"gin-expenseapp-api/config"
	"gin-expenseapp-api/model"
	"gin-expenseapp-api/library"

	"github.com/gin-gonic/gin"
)

// TransactionRequest Structure of the API request parameters
type TransactionRequest struct {
	// Description Description of the transaction
	Description string `json:"description"`
	// Amount Transaction amount
	Amount float64 `json:"amount"`
}

// Credit Controller for the credit transaction
func Credit(c *gin.Context) {
	var request TransactionRequest
	c.BindJSON(&request)
	cookie, _ := c.Request.Cookie("token")
	token := cookie.Value

	claims := library.GetClaims(token, config.Auth_secret)

	response := model.AddTransaction(claims.UserID, request.Description, request.Amount, 0, 0)

	c.JSON(200, response)
}
