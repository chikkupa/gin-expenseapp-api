package model

import (
	"github.com/gin-gonic/gin"
)

const (
	//Success Status: Success
	Success = 0 
	//Failed Status: Failed
	Failed = 1
	//NotFound Status: Not Found
	NotFound = 2
	//Unauthorised Staus: Unauthorised
	Unauthorised = 3
	//DatabaseConnectionError Status: Database connection error
	DatabaseConnectionError = 4
	//DatabaseError Status: Database error
	DatabaseError = 5
)

var responseCodes = []struct {
	Status int 
	Result string
}{
	{ Status: 200, Result: "Success" },
	{ Status: 200, Result: "Failed"},
	{ Status: 404, Result: "NotFound"},
	{ Status: 401, Result: "UnAuthorised"},
	{ Status: 200, Result: "DatabaseConnectionError"},
	{ Status: 200, Result: "DatabaseError"},
}
//Status List of response status
var Status = []int{200, 200, 404, 401, 200, 200}
//Results List of response results
var Results = []string{"Success", "Failed", "NotFound", "UnAuthorised", "DatabaseConnectionError", "DatabaseError"}

//Response Format of the response
func Response(result int, message string, data interface{}) gin.H {
	if(data != nil){
		return gin.H{
			"status" : responseCodes[result].Status,
			"result" : responseCodes[result].Result,
			"message" : message,
			"data" : data,
		}
	}
	return gin.H{
		"status" : responseCodes[result].Status,
		"result" : responseCodes[result].Result,
		"message" : message,
	}
}