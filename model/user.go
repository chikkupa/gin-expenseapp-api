package model

import (
	"database/sql"
	"gin-expenseapp-api/config"
	"gin-expenseapp-api/library"
)

// IsAuthenticatedUser Checks whether the username and password match
func IsAuthenticatedUser(username string, password string) (bool, error) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		return false, err
	}
	defer db.Close()

	password = library.GetMd5(password)
	
	// Execute the query
    results, err := db.Query("SELECT id FROM user WHERE username=? AND password=?", username, password)
    if err != nil {
        return false, err
	}
	
	if results.Next() {
		return true, nil
	}

	return false, nil
}