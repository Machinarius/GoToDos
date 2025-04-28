package routes

import (
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine) {
	r.GET("/api/todos", GetToDosHandler)
}
