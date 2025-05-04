package main

import (
	"log"
	"os"

	"wallet-topup/internal/api/wallet/handlers"
	"wallet-topup/internal/api/wallet/repo"
	"wallet-topup/internal/api/wallet/routes"
	"wallet-topup/internal/api/wallet/services"
	"wallet-topup/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load("config/.env"); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize database
	db := database.InitDB()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error getting database instance: %v", err)
			return
		}
		sqlDB.Close()
	}()

	// Initialize wallet module
	walletRepo := repo.NewRepository(db)
	walletService := services.NewService(walletRepo)
	walletHandler := handlers.NewHandler(walletService)

	// Initialize router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r, walletHandler)

	// Basic health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
