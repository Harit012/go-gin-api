package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-gin-api/models"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_NAME", "mydb"),
		getEnv("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 👇 This creates the users table if it doesn't exist
    err = db.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Failed to auto-migrate:", err)
    }

	DB = db
	log.Println("📦 Connected to PostgreSQL")
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
