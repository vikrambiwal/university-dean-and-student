package controllers

import (
	"fmt"
	"net/http"
	"time"
	"university-dean-and-student/app/models"
	"university-dean-and-student/app/utility"

	"github.com/gin-gonic/gin"
)

func Session(c *gin.Context) {
	user := CurrentUser(c)

	var requestBody models.Session
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		requestBody.StudentId = int64(user.ID)
		// validateSessionTime()
		fmt.Println("requestBody--", requestBody)
		err := requestBody.ValidateSession()
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid request body", "error": err})
			return
		}

		result := utility.Database().Create(&requestBody)
		if result.Error != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Error in signup", "error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Signup successful", "user": requestBody})
		return
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid request body"})
	}

	c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid user type"})
}

func Sessions(c *gin.Context) {
	user := CurrentUser(c)

	if user.UserType == "STUDENT" {
		sessions := models.SessionListForStudents(int64(user.ID))
		c.JSON(http.StatusOK, gin.H{"message": "Successful", "sessions": sessions})
		return
	} else if user.UserType == "DEAN" {
		sessions := models.SessionListForDean(int64(user.ID))
		c.JSON(http.StatusOK, gin.H{"message": "Successful", "sessions": sessions})
		return
	}

	c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid user type"})
}

func validateSessionTime(c *gin.Context, date string) *time.Time {
	sessionDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid session date"})
		return nil
	}

	return &sessionDate
}
