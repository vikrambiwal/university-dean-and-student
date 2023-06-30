package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Sessions(c *gin.Context) {
	// var requestBody models.User
	// if err := c.ShouldBindJSON(&requestBody); err != nil {
	// 	err := requestBody.ValidateUser()
	// 	if err != nil {
	// 		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid request body", "error": err})
	// 		return
	// 	}
	// 	requestBody.Token = uuid.Must(uuid.NewRandom()).String()

	// 	result := utility.Database().Create(&requestBody)
	// 	if result.Error != nil {
	// 		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Error in signup", "error": result.Error.Error()})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{"message": "Signup successful", "user": requestBody})
	// 	return
	// } else {
	// }
	c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid request body"})
}
