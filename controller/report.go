package controller

import (
	"time"
	"gin-expenseapp-api/model"
	"gin-expenseapp-api/config"
	"gin-expenseapp-api/library"
	
	"github.com/gin-gonic/gin"
)

// ReportRequest Structure for API request
type ReportRequest struct {
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
}

// GetTransactionReport Get transaction reports for a given interval
func GetTransactionReport(c *gin.Context) {
	var request ReportRequest
	err := c.BindJSON(&request)

	if err != nil{
		c.AbortWithStatusJSON(400, model.Response(model.BadRequestError, err.Error(), nil))
		return
	}

	sStartDate := ""
	sEndDate := ""

	cookie, _ := c.Request.Cookie("token")
	token := cookie.Value

	layout := "2006-01-02"
	if request.StartDate != "" {
		startDate, err := time.Parse(layout, request.StartDate)

		if err != nil{
			c.JSON(400, model.Response(model.BadRequestError, "Start Date: Invalid date format. Required YYYY-mm-dd", nil))
			return
		}
		sStartDate = startDate.Format(layout)
	}
	
	if request.EndDate != "" {
		endDate, err := time.Parse(layout, request.EndDate)

		if err != nil{
			c.JSON(400, model.Response(model.BadRequestError, "End Date: Invalid date format. Required YYYY-mm-dd", nil))
			return
		}

		sEndDate = endDate.Format(layout)
	}

	claims := library.GetClaims(token, config.Auth_secret)

	response := model.GetReport(claims.UserID, sStartDate, sEndDate)

	c.JSON(200, response)
}