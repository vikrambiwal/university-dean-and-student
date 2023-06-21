package main

import (
	"university-dean-and-student/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/login", controllers.Login)

	router.Run("localhost:8080")
}
