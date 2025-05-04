package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"machinarius.github.io/gotodos/internal/api/routes"
	"machinarius.github.io/gotodos/internal/store/database"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file", envErr)
	}

	database.Initialize()

	r := gin.Default()
	routes.ConfigureRoutes(r)
	r.Run("localhost:8080")
}
