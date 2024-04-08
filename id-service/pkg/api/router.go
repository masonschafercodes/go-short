package api

import (
	"github.com/gin-gonic/gin"
	"github.com/masonschafercodes/go-short/id-service/pkg/api/ids"
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

	v1RouteGroup := r.Group("/api/v1")

	v1RouteGroup.GET("/metrics", prometheusHandler())
	v1RouteGroup.GET("/id", ids.CreateShortLinkId)
	return r
}
