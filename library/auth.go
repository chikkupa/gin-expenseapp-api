package library

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	UserID int `json:"userid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GetClaims Get claims from jwt token
func GetClaims(key string, authSecret string) Claims {
	// Initialize a new instance of `Claims`
	claims := &Claims{}

	jwt.ParseWithClaims(key, claims, func(token *jwt.Token) (interface{}, error) {
		return authSecret, nil
	})

	return *claims
}