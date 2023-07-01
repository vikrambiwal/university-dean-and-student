package routes

import (
	"net/http"
	"strings"
	"university-dean-and-student/app/models"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")

		authFields := strings.Fields(auth)
		if len(authFields) != 2 || strings.ToLower(authFields[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid token."})
			return
		}

		user := models.AuthUser(authFields[1])
		if user == nil || user.Token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid token."})
			return
		}

		c.Next()
	}
}
