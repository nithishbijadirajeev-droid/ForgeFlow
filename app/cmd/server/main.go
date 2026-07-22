// ForgeFlow API
//
// @title ForgeFlow API
// @version 1.0
// @description Production-ready backend built with Go, Gin, PostgreSQL, JWT Authentication, Docker, and GORM.
//
// @contact.name Nithish Bijadi Rajeev
// @contact.email nithishcsu@gmail.com
//
// @license.name MIT
//
// @host localhost:8080
// @BasePath /
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your JWT token in the format: Bearer <your_token>

package main

import (
	_ "github.com/nithishbijadirajeev-droid/forgeflow/docs"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/config"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/database"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/logger"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/models"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/routes"

	"go.uber.org/zap"
)

func main() {

	// Initialize logger
	log := logger.GetLogger()
	defer logger.Sync()

	// Load environment variables
	config.LoadEnv()

	// Connect to database
	database.Connect()

	// Run database migrations
	if err := database.DB.AutoMigrate(
		&models.User{},
		&models.Project{},
	); err != nil {

		log.Fatal(
			"Database migration failed",
			zap.Error(err),
		)
	}

	// Setup routes
	router := routes.SetupRouter()

	// Get server port
	port := config.Get("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// Startup log
	log.Info(
		"ForgeFlow server starting",
		zap.String("port", port),
	)

	// Start server
	if err := router.Run(":" + port); err != nil {

		log.Fatal(
			"Failed to start server",
			zap.Error(err),
		)
	}
}
