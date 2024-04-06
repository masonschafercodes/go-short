package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/masonschafercodes/go-short/pkg/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(gin.DebugMode)

	r := api.InitRouter()

	if err := r.Run(":3003"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
