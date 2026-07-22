package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/handlers"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/middleware"
)

func SetupRouter() *gin.Engine {

	router := gin.New()

	router.Use(gin.Recovery())

	// Order matters
	router.Use(middleware.RequestID())
	router.Use(middleware.RequestLogger())
	router.Use(middleware.ErrorHandler())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		auth := handlers.NewAuthHandler()
		project := handlers.NewProjectHandler()

		api.POST("/register", auth.Register)
		api.POST("/login", auth.Login)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())

		protected.POST("/projects", project.Create)
		protected.GET("/projects", project.GetAll)
		protected.GET("/projects/:id", project.GetByID)
		protected.PUT("/projects/:id", project.Update)
		protected.DELETE("/projects/:id", project.Delete)
	}

	return router
}
