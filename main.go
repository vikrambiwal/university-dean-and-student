package main

import (
	"log"
	"university-dean-and-student/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	router := gin.Default()
	public := router.Group("/api")

	public.POST("/login", controllers.Login)
	public.POST("/register", controllers.Register)

	router.Run("localhost:8080")
}
