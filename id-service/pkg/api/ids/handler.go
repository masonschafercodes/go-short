package ids

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func generateID(url string) string {
	rand.Seed(time.Now().UnixNano())
	randomURL := fmt.Sprintf("%s%d", url, rand.Intn(100000))
	hasher := sha256.New()
	hasher.Write([]byte(randomURL))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)[:5]
}

func CreateShortLinkId(ctx *gin.Context) {
	linkQuery := ctx.Query("link")

	if linkQuery == "" {
		ctx.JSON(400, gin.H{
			"error": "link is required",
		})
		return
	}

	id := generateID(linkQuery)

	ctx.JSON(200, gin.H{
		"id": id,
	})
}
