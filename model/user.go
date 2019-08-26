package model

import (
	"database/sql"
	"gin-expenseapp-api/config"
	"gin-expenseapp-api/library"
)

// User User structure
type User struct {
	// UserId User Id
	UserID int
	Username string
	Balance float64
}

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

// GetUserDetails Get user details
func GetUserDetails(username string) (User, error) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	var user User

	if err != nil {
		return user, err
	}
	defer db.Close()

	result, err := db.Query("SELECT id, username, balance FROM user WHERE username=?", username)
    if err != nil {
        return user, err
	}
	
	if result.Next() {
		err = result.Scan(&user.UserID, &user.Username, &user.Balance)
		if err != nil {
			return user, err
		}

		return user, nil
	}

	return user, nil
}

// GetUserDetailsFromID Get user details
func GetUserDetailsFromID(userID int) (User, error) {
	db, err := sql.Open(config.Mysql, config.Dbconnection)
	var user User

	if err != nil {
		return user, err
	}
	defer db.Close()

	result, err := db.Query("SELECT id, username, balance FROM user WHERE id=?", userID)
    if err != nil {
        return user, err
	}
	
	if result.Next() {
		err = result.Scan(&user.UserID, &user.Username, &user.Balance)
		if err != nil {
			return user, err
		}

		return user, nil
	}

	return user, nil
}

// UpdateBalance Update the balance of user
func UpdateBalance(userID int, balance float64) error {
	db, err := sql.Open(config.Mysql, config.Dbconnection)

	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Query("UPDATE user set balance=? WHERE id=?", balance, userID)

	if err != nil {
		return err
	}

	return nil
}