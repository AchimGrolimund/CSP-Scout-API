package main

import (
	"log"
	"os"


	"github.com/AchimGrolimund/CSP-Scout-API/pkg/application"
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/infrastructure/mongodb"
	"github.com/AchimGrolimund/CSP-Scout-API/pkg/interfaces/http/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
)

var (
	Build = "raw"
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

	// Create service and handler
	service := application.NewReportService(repo)
	handler := handlers.NewReportHandler(service)

	// Initialize Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.Default())

	// Setup routes
	handlers.SetupRoutes(router, handler)

	// Start the server
	router.Run()
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
