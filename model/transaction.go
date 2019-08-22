package model

import (
	"database/sql"
	"gin-expenseapp-api/config"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// AddTransaction Add a new transaction
func AddTransaction(description string, credit float64, debit float64, balance float64) gin.H {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		panic(err.Error())

		return gin.H{
			"status":  200,
			"result":  "Error",
			"message": err.Error(),
		}
	}
	defer db.Close()

	_, err = db.Query("INSERT INTO transaction (description, debit, credit, balance) VALUES (?, ?, ?, ?)", description, debit, credit, balance)

	if err != nil {
		panic(err.Error())
		return gin.H{
			"status":  200,
			"result":  "Error",
			"message": err.Error(),
		}
	}

	return gin.H{
		"status":  200,
		"result":  "Success",
		"message": "Transaction completed successfully!",
	}
}
