package controller

import (
	"github.com/gin-gonic/gin"
	"gin-expenseapp/model"
)

type TransactionRequest struct { 
	Description string `json: "description"`
	Amount float64	`json: "amount:`
}

func Credit(c *gin.Context){
	var request TransactionRequest
	c.BindJSON(&request)
	
	response := model.AddTransaction(request.Description, 0, request.Amount, 0)
	
	c.JSON(200, response)
}