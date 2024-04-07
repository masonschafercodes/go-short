package api

import (
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/masonschafercodes/go-short/pkg/api/healthcheck"
	"github.com/masonschafercodes/go-short/pkg/api/links"
	"github.com/masonschafercodes/go-short/pkg/db"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())

	store := ratelimit.RedisStore(&ratelimit.RedisOptions{
		RedisClient: db.GetRedisClient(),
		Rate:        time.Second,
		Limit:       5,
	})

	rateLimitingMiddleware := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	v1RouteGroup := r.Group("/api/v1")

	v1RouteGroup.GET("/", healthcheck.HandleHealthCheck)

	v1RouteGroup.POST("/link", rateLimitingMiddleware, links.CreateShortLink)

	return r
}
