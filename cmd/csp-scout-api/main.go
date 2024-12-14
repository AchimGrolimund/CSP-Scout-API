package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AchimGrolimund/CSP-Scout-API/pkg/application"
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/infrastructure/mongodb"
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/interfaces/http/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	Build   = "raw"
	Version = "raw"
)

func main() {
	// Load environment variables
	if err := godotenv.Load("configs/.local.env"); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// MongoDB configuration
	mongoURI := getEnv("MONGODB_URI", "mongodb://localhost:27017")
	dbName := getEnv("MONGODB_DATABASE", "csp_scout")
	collectionName := getEnv("MONGODB_COLLECTION", "reports")

	// Initialize MongoDB repository
	repo, err := mongodb.NewMongoRepository(mongoURI, dbName, collectionName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Create service
	service := application.NewService(repo)

	// Initialize Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.Default())

	// Setup routes
	handlers.RegisterRoutes(router, service)

	// Get server port from environment variables
	port := getEnv("SERVER_PORT", "8080")

	// Start the server with configured port
	log.Printf("Server starting on port %s", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
