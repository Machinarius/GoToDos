package main

import (
	"github.com/gin-gonic/gin"
	"machinarius.github.io/gotodos/internal/api/routes"
)

func main() {
	r := gin.Default()
	routes.ConfigureRoutes(r)
	r.Run("localhost:8080")
}
