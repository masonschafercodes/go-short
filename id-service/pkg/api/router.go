package api

import (
	"github.com/gin-gonic/gin"
	"github.com/masonschafercodes/go-short/id-service/pkg/api/ids"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())

	v1RouteGroup := r.Group("/api/v1")

	v1RouteGroup.GET("/id", ids.CreateShortLinkId)
	return r
}
