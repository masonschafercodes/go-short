package api

import (
	"github.com/gin-gonic/gin"
	"github.com/masonschafercodes/go-short/redirection-service/pkg/api/links"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())

	r.GET("/:id", links.RedirectToLink)
	return r
}
