package database

import (
	"fmt"
	"log"
	"os"

	"wallet-topup/internal/database/handlers"
	"wallet-topup/internal/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() handlers.HealthStats

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	// DB returns the underlying *gorm.DB instance
	DB() *gorm.DB
}

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(&entities.User{}, &entities.Transaction{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Create a test user if none exists
	var count int64
	db.Model(&entities.User{}).Count(&count)
	if count == 0 {
		testUser := &entities.User{
			ID:      1,
			Balance: 0,
		}
		if err := db.Create(testUser).Error; err != nil {
			log.Printf("Failed to create test user: %v", err)
		}
	}

	return db
}
