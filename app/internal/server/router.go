package server

import (
	"github.com/gin-gonic/gin"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/handlers"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", handlers.Health)
	router.GET("/ready", handlers.Ready)
	router.GET("/version", handlers.Version)

	return router
}
