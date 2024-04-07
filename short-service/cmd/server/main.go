package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/masonschafercodes/go-short/pkg/api"
)

func main() {
	gin.SetMode(gin.DebugMode)

	r := api.InitRouter()

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
