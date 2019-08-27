package model

import (
	"database/sql"
	"gin-expenseapp-api/config"

	"github.com/gin-gonic/gin"
)

// AddTransaction Add a new transaction
func AddTransaction(userID int, date string, description string, credit float64, debit float64, balance float64) gin.H {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		return Response(DatabaseConnectionError, err.Error(), nil)
	}
	defer db.Close()

	user, err := GetUserDetailsFromID(userID)

	if err != nil {
		return Response(DatabaseError, err.Error(), nil)
	}

	userBalance := user.Balance

	userBalance += credit
	userBalance -= debit

	UpdateBalance(userID, userBalance)

	_, err = db.Query("INSERT INTO transaction (user_id, date, description, debit, credit, balance) VALUES (?, ?, ?, ?, ?, ?)", userID, date, description, debit, credit, userBalance)

	if err != nil {
		return Response(DatabaseError, err.Error(), nil)
	}

	return Response(Success, "Transaction completed successfully!", nil)
}
