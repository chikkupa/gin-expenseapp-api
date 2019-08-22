package controller

import (
	"gin-expenseapp-api/model"

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

	response := model.AddTransaction(request.Description, 0, request.Amount, 0)

	c.JSON(200, response)
}
