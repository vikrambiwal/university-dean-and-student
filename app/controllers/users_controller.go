package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"university-dean-and-student/app/models"
	"university-dean-and-student/app/utility"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CurrentUser(c *gin.Context) *models.User {
	auth := c.Request.Header.Get("Authorization")

	authFields := strings.Fields(auth)
	if len(authFields) != 2 || strings.ToLower(authFields[0]) != "bearer" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid token."})
		return nil
	}

	return models.AuthUser(authFields[1])
}

func Register(c *gin.Context) {
	var requestBody models.User
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		err := requestBody.ValidateUser()
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid request body", "error": err})
			return
		}
		requestBody.Token = uuid.Must(uuid.NewRandom()).String()
		requestBody.CreatedAt = time.Now()
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

func DeanList(c *gin.Context) {
	users := models.DeanList()
	c.JSON(http.StatusOK, gin.H{"message": "Successful", "users": users})
}
