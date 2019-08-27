package model

import (
	"database/sql"
	"gin-expenseapp-api/config"

	"github.com/gin-gonic/gin"
)

// Transaction Structure for transaction item
type Transaction struct {
	ID int `json:"id"`
	UserID int `json:"userid"`
	Date string `json:"date"`
	Description string `json:"description"`
	Debit float64 `json:"debit"`
	Credit float64 `json:"credit"`
	Balance float64 `json:"balance"`
	Time string `json:"Time"`
}

// GetReport To get transaction report
func GetReport(userID int, startDate string, endDate string) gin.H {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		return Response(DatabaseConnectionError, err.Error(), nil)
	}
	defer db.Close()

	query := "SELECT id, user_id, date, description, debit, credit, balance, time FROM transaction WHERE user_id=?"
	if startDate != "" {
		query += " AND date>='" + startDate + "'"
	}
	if endDate != "" {
		query += " AND date<='" + endDate + "'"
	}
	query += " ORDER BY id DESC"
	results, err := db.Query(query, userID)
	
	if err != nil {
		return Response(DatabaseError, err.Error(), nil)
	}

	var items []Transaction

	for results.Next() {
        var item Transaction
        // for each row, scan the result into our tag composite object
        err = results.Scan(&item.ID, &item.UserID, &item.Date, &item.Description, &item.Debit, &item.Credit, &item.Balance, &item.Time)
        if err != nil {
            return Response(DatabaseError, err.Error(), nil)
		}
		
		items = append(items, item)
	}
	
	return Response(Success, "", map[string]interface{}{
		"count" : len(items),
		"entries" : items,
	})
}