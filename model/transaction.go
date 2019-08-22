package model

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func AddTransaction(description string, credit float64, debit float64, balance float64) gin.H{
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/expenseapp")

    if err != nil {
		panic(err.Error())
		
		return gin.H{
			"status" : 200,
			"result" : "Error",
			"message": err.Error(),
		}
	}
	defer db.Close()

	_, err = db.Query("INSERT INTO transaction (description, debit, credit, balance) VALUES (?, ?, ?, ?)", description, debit, credit, balance);

    if err != nil {
		panic(err.Error())
		return gin.H{
			"status" : 200,
			"result" : "Error",
			"message": err.Error(),
		}
	}

	return gin.H{
		"status" : 200,
		"result" : "Success",
		"message": "Transaction completed successfully!",
	}
}