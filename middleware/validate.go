package middleware

import (
	"io/ioutil"
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-expenseapp-api/library"
	"gin-expenseapp-api/model"
)
// ValidateRequest Validate API request body
func ValidateRequest(c *gin.Context){
	if c.Request.Method != http.MethodGet {
		b, _ := ioutil.ReadAll(c.Request.Body)

		if(!library.IsJSON(string(b))){
			c.AbortWithStatusJSON(400, model.Response(model.BadRequestError, "Invalid JSON Request", nil))
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	}

	c.Next();
}