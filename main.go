package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("❌ Error loading .env file")
	}
}

func main() {
	LoadEnv()
	ConnectDatabase()
	r := gin.Default()
	SetupRoutes(r)
	r.Run(":8080")
}
