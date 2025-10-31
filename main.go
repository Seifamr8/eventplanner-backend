package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Warning: .env file not found, using system environment variables")
	}
}

func main() {
	LoadEnv()
	ConnectDatabase()

	r := gin.Default()

	// Enable CORS
	r.Use(CORSMiddleware())

	// Register routes
	SetupRoutes(r)

	// Start server
	log.Println("🚀 Server running on http://localhost:8080")
	r.Run(":8080")
}
