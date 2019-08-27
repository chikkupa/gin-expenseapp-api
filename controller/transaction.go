package controller

import (
	"time"
	"gin-expenseapp-api/config"
	"gin-expenseapp-api/model"
	"gin-expenseapp-api/library"

	"github.com/gin-gonic/gin"
)

// TransactionRequest Structure of the API request parameters
type TransactionRequest struct {
	// Date Date of transaction
	Date string `json:"date"`
	// Description Description of the transaction
	Description string `json:"description"`
	// Amount Transaction amount
	Amount float64 `json:"amount"`
}

// Credit Controller for the credit transaction
func Credit(c *gin.Context) {
	var request TransactionRequest
	err := c.BindJSON(&request)
	if err != nil{
		c.AbortWithStatusJSON(400, model.Response(model.BadRequestError, err.Error(), nil))
		return
	}

	cookie, _ := c.Request.Cookie("token")
	token := cookie.Value

	layout := "2006-01-02"
	date, err := time.Parse(layout, request.Date)

	if err != nil{
		c.JSON(400, model.Response(model.BadRequestError, "Invalid date format. Required YYYY-mm-dd", nil))
		return
	}

	claims := library.GetClaims(token, config.Auth_secret)

	response := model.AddTransaction(claims.UserID, date.Format(layout), request.Description, request.Amount, 0, 0)

	c.JSON(200, response)
}

// Debit Controller for the debit transaction
func Debit(c *gin.Context) {
	var request TransactionRequest
	err := c.BindJSON(&request)

	if err != nil{
		c.AbortWithStatusJSON(400, model.Response(model.BadRequestError, err.Error(), nil))
		return
	}

	cookie, _ := c.Request.Cookie("token")
	token := cookie.Value

	layout := "2006-01-02"
	date, err := time.Parse(layout, request.Date)

	if err != nil{
		c.JSON(400, model.Response(model.BadRequestError, "Invalid date format. Required YYYY-mm-dd", nil))
		return
	}

	claims := library.GetClaims(token, config.Auth_secret)

	response := model.AddTransaction(claims.UserID, date.Format(layout), request.Description, 0, request.Amount, 0)

	c.JSON(200, response)
}