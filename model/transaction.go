package model

import (
	"database/sql"
	"gin-expenseapp-api/config"

	"github.com/gin-gonic/gin"
)

// AddTransaction Add a new transaction
func AddTransaction(description string, credit float64, debit float64, balance float64) gin.H {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		return Response(DatabaseConnectionError, err.Error(), nil)
	}
	defer db.Close()

	_, err = db.Query("INSERT INTO transaction (description, debit, credit, balance) VALUES (?, ?, ?, ?)", description, debit, credit, balance)

	if err != nil {
		return Response(DatabaseError, err.Error(), nil)
	}

	return Response(Success, "Transaction completed successfully!", nil)
}
