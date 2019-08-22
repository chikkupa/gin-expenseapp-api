package controller

import (
	"github.com/gin-gonic/gin"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type TransactionRequest struct { 
	Description string `json: "description"`
	Amount float64	`json: "amount:`
}

func Credit(c *gin.Context){
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/expenseapp")

    if err != nil {
		c.JSON(200, gin.H{
			"status" : 200,
			"result" : "Error",
			"message": err.Error(),
		})
		panic(err.Error())
		return
	}
	defer db.Close()

	var request TransactionRequest
	c.BindJSON(&request)
	
	_, err = db.Query("INSERT INTO transaction (description, debit, credit, balance) VALUES (?, ?, ?, ?)", request.Description, 0, request.Amount, 0);

    if err != nil {
        c.JSON(200, gin.H{
			"status" : 200,
			"result" : "Error",
			"message": err.Error(),
		})
		panic(err.Error())
		return
	}

	c.JSON(200, gin.H{
		"description" : request.Description,
		"amount" : request.Amount,
	})
}