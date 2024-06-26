package api

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/masonschafercodes/go-short/redirection-service/pkg/api/healthcheck"
	"github.com/masonschafercodes/go-short/redirection-service/pkg/api/links"
	"github.com/masonschafercodes/go-short/redirection-service/pkg/db"
	"github.com/masonschafercodes/go-short/redirection-service/pkg/worker"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())

	dbClient := db.GetConnection()
	numOfWorkers := runtime.NumCPU() * 2

	for i := 0; i < numOfWorkers; i++ {
		go worker.UpdateWorker(dbClient)
	}

	r.GET("/metrics", prometheusHandler())
	r.GET("/health", healthcheck.HandleHealthCheck)
	r.GET("/:id", links.RedirectToLink)
	return r
}
