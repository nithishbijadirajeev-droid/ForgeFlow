package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/handlers"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/middleware"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	projectHandler := handlers.NewProjectHandler()
	authHandler := handlers.NewAuthHandler()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	router.GET("/ready", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "READY",
		})
	})

	router.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "v1.0.0",
		})
	})

	api := router.Group("/api/v1")

	// Public routes
	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.POST("/projects", projectHandler.Create)
	protected.GET("/projects", projectHandler.GetAll)
	protected.GET("/projects/:id", projectHandler.GetByID)
	protected.PUT("/projects/:id", projectHandler.Update)
	protected.DELETE("/projects/:id", projectHandler.Delete)

	return router
}