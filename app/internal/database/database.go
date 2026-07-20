package database

import (
	"fmt"
	"log"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	host := config.Get("DB_HOST")
	port := config.Get("DB_PORT")
	user := config.Get("DB_USER")
	password := config.Get("DB_PASSWORD")
	dbname := config.Get("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	log.Println("✅ Connected to PostgreSQL")
}