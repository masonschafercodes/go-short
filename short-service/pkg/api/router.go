package api

import (
	"github.com/gin-gonic/gin"
	"github.com/masonschafercodes/go-short/pkg/api/healthcheck"
	"github.com/masonschafercodes/go-short/pkg/api/links"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())

	v1RouteGroup := r.Group("/api/v1")

	v1RouteGroup.GET("/", healthcheck.HandleHealthCheck)

	v1RouteGroup.POST("/link", links.CreateShortLink)

	return r
}
