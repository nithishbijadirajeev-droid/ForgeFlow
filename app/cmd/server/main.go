package main

import (
	"log"
	"os"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/database"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/models"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/routes"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/config"
)

func main() {
	config.LoadEnv()

	if os.Getenv("SERVER_PORT") == "" {
		os.Setenv("SERVER_PORT", "8080")
	}

	database.Connect()

	err := database.DB.AutoMigrate(
    &models.User{},
    &models.Project{},
)

	if err != nil {
		log.Fatal(err)
	}

	router := routes.SetupRouter()

	log.Println("🚀 ForgeFlow started on :8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}