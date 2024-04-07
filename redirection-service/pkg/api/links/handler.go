package links

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masonschafercodes/go-short/redirection-service/pkg/db"
)

type MessageFromIdService struct {
	ID string `json:"id"`
}

func RedirectToLink(ctx *gin.Context) {
	shortId := ctx.Param("id")

	if shortId == "" {
		log.Println("No short id provided")
		ctx.Status(404)
		return
	}

	client := db.GetConnection()

	err := client.Ping(ctx)
	if err != nil {
		log.Println("Error pinging database", err)
		ctx.Status(500)
		return
	}

	var originalURL string
	err = client.QueryRow(ctx, "SELECT original_url FROM links WHERE short_url = ($1)", shortId).Scan(&originalURL)
	if err != nil {
		log.Println("Error querying database", err.Error())
		ctx.Status(404)
		return
	}

	_, err = client.Exec(ctx, "UPDATE links SET access_count = access_count + 1 WHERE short_url = ($1)", shortId)

	if err != nil {
		log.Println("Error updating access count", err)
		ctx.Status(500)
		return
	}

	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.Redirect(http.StatusMovedPermanently, originalURL)
}
