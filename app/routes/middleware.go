package routes

import (
	"net/http"
	"university-dean-and-student/app/models"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")

		if auth == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid token."})
			return
		}

		user := models.AuthUser(auth)
		if user == nil || user.Token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid token."})
			return
		}

		c.Next()
	}
}
