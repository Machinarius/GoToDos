package routes

import (
	"net/http"

	"machinarius.github.io/gotodos/internal/api"
	"machinarius.github.io/gotodos/internal/store"

	"github.com/gin-gonic/gin"
)

func getToDosHandler(c *gin.Context) {
	c.JSON(http.StatusOK, store.GetAll())
}

func createToDosHandler(c *gin.Context) {
	var req api.CreateToDoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, store.Create(req.Text))
}
