package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/masonschafercodes/go-short/id-service/pkg/api"
)

func main() {
	gin.SetMode(gin.DebugMode)

	r := api.InitRouter()

	if err := r.Run(":3004"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
