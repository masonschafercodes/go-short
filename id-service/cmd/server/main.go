package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/masonschafercodes/go-short/id-service/pkg/api"
)

func main() {
	gin.SetMode(gin.DebugMode)

	r := api.InitRouter()

	if err := r.Run(":" + os.Getenv("PORT")); err != nil { // TODO: get port from env
		log.Fatalf("Failed to start server: %v", err)
	}
}
