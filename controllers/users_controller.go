package controllers

import (
	"fmt"
	"net/http"
	"university-dean-and-student/models"
	"university-dean-and-student/utility"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Register(c *gin.Context) {
	var requestBody models.User
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		err := requestBody.ValidateUser()
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid request body", "error": err})
			return
		}
		requestBody.Token = uuid.Must(uuid.NewRandom()).String()

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
}

func Login(c *gin.Context) {
	var requestBody models.User
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		user := models.FindUser(requestBody.UserName, requestBody.Password)
		fmt.Println("zzzzz", user)
		if user == nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid login details."})
			return
		}

		user.Token = uuid.Must(uuid.NewRandom()).String()
		result := utility.Database().Save(&user) //.Where("id = ?", requestBody.Id)
		if result.Error != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Error in login", "error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}
