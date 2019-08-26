package middleware

import (
	"net/http"
	"gin-expenseapp-api/model"
	"gin-expenseapp-api/library"
	"gin-expenseapp-api/config"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(config.Auth_secret)

// AuthRequest Structure for authentication request
type AuthRequest struct {
	Token string `json:"token"`
}

// Authenticate Middleware function to authenticate API
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("token")
		
		if err != nil {
			if err == http.ErrNoCookie {
				c.AbortWithStatusJSON(400, model.Response(model.BadRequestError, "No token found", nil))
				return
			}
		}
		token := cookie.Value

		if token == "" {
			c.AbortWithStatusJSON(401, model.Response(model.Unauthorised, "Unauthorised request", nil))
			return
		}

		// Initialize a new instance of `Claims`
		claims := &library.Claims{}

		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(401, model.Response(model.Unauthorised, "Invalid Signature: Unauthorised request", nil))
				return
			}
			c.AbortWithStatusJSON(400, model.Response(model.BadRequestError, "Token Error: Bad Request", nil))
			return
		}

		if !tkn.Valid {
			c.AbortWithStatusJSON(401, model.Response(model.Unauthorised, "Unauthorised request", nil))
			return
		}
		
		c.Next()
	}
}