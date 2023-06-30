package main

import (
	"log"

	"university-dean-and-student/app/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	routes.SetupRoutes()
}
