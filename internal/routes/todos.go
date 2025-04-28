package routes

import (
	"net/http"

	"machinarius.github.io/gotodos/internal/store"

	"github.com/gin-gonic/gin"
)

func getToDosHandler(c *gin.Context) {
	c.JSON(http.StatusOK, store.GetAll())
}
