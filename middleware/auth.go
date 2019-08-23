package middleware

import (
	"fmt"
	"gin-expenseapp-api/model"

	"github.com/gin-gonic/gin"
)

//AuthRequest Structure for authentication request
type AuthRequest struct {
	Token string `json:"token"`
}

//Authenticate Middleware function to authenticate API
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		fmt.Println(token)
		if token == "" {
			respondWithError(c, 401, "Unauthorised request")
			return
		}

		c.Next()
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, model.Response(model.Unauthorised, message.(string), nil))
}