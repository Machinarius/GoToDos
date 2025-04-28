package routes

import (
	"net/http"

	"machinarius.github.io/gotodos/internal/store"

	"github.com/gin-gonic/gin"
)

func GetToDosHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, store.GetAll())
}
