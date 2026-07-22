package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "v0.1.0",
	})
}
