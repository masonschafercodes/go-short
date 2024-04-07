package links

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masonschafercodes/go-short/redirection-service/pkg/db"
	"github.com/redis/go-redis/v9"
)

type MessageFromIdService struct {
	ID string `json:"id"`
}

func updateCountForShortId(client *pgxpool.Pool, shortId string, ctx *gin.Context) error {
	_, err := client.Exec(ctx, "UPDATE links SET access_count = access_count + 1 WHERE short_url = ($1)", shortId)

	if err != nil {
		log.Println("Error updating access count", err)
		return err
	}

	return nil
}

func handleRedirect(ctx *gin.Context, originalURL string) {
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.Redirect(http.StatusMovedPermanently, originalURL)
}

func RedirectToLink(ctx *gin.Context) {
	shortId := ctx.Param("id")

	if shortId == "" {
		log.Println("No short id provided")
		ctx.Status(404)
		return
	}

	rdb := db.GetRedisClient()
	val, err := rdb.Get(ctx, shortId).Result()

	switch {

	case err == redis.Nil:
		log.Println("Value not found in redis, querying database")
		client := db.GetConnection()

		var originalURL string
		err = client.QueryRow(ctx, "SELECT original_url FROM links WHERE short_url = ($1)", shortId).Scan(&originalURL)

		if err != nil {
			log.Println("Error querying database", err.Error())
			ctx.Status(404)
			return
		}

		rdb.Set(ctx, shortId, originalURL, time.Minute*5)
		err = updateCountForShortId(client, shortId, ctx)

		if err != nil {
			log.Println("Error updating access count", err)
			ctx.Status(500)
			return
		}

		handleRedirect(ctx, originalURL)

	case err != nil:
		log.Println("Error querying redis", err)
		ctx.Status(500)

	case val == "":
		log.Println("Value is empty")
		ctx.Status(404)

	case val != "":
		log.Println("Value found in redis")

		client := db.GetConnection()

		err := updateCountForShortId(client, shortId, ctx)

		if err != nil {
			log.Println("Error updating access count", err)
			ctx.Status(500)
			return
		}

		handleRedirect(ctx, val)
	}
}
