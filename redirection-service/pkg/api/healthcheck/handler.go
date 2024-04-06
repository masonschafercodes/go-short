package healthcheck

import "github.com/gin-gonic/gin"

func HandleHealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
