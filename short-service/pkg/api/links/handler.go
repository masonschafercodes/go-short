package links

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/masonschafercodes/go-short/pkg/db"
)

type MessageFromIdService struct {
	ID string `json:"id"`
}

func isValidHTTPSUrl(url string) bool {
	// Regular expression for validating an HTTPS URL
	regex := `^https:\/\/[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,3}(:[a-zA-Z0-9]*)?\/?([a-zA-Z0-9\-\/]*)?$`
	re := regexp.MustCompile(regex)
	return re.MatchString(url)
}

func CreateShortLink(ctx *gin.Context) {
	linkQuery := ctx.Query("link")

	if linkQuery == "" {
		log.Println("No link provided")
		ctx.JSON(400, gin.H{
			"error": "link is required",
		})
		return
	}

	if !isValidHTTPSUrl(linkQuery) {
		log.Println("Invalid URL provided")
		ctx.JSON(400, gin.H{
			"error": "link must be a valid HTTPS URL",
		})
		return
	}

	client := db.GetConnection()

	err := client.Ping(ctx)
	if err != nil {
		log.Println("Error pinging database", err)
		ctx.JSON(500, gin.H{
			"error": http.StatusText(500),
		})
		return
	}

	idServiceUrl := os.Getenv("ID_SERVICE_URL")
	if idServiceUrl == "" {
		log.Println("ID_SERVICE_URL is not set")
		ctx.JSON(500, gin.H{
			"error": http.StatusText(500),
		})
		return
	}
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/id?link=%s", idServiceUrl, linkQuery))
	if err != nil {
		log.Println("Error getting id from id service", err)
		ctx.JSON(500, gin.H{
			"error": http.StatusText(500),
		})
		return
	}

	if resp.StatusCode != 200 {
		log.Println("Error getting id from id service", resp.StatusCode)
		ctx.JSON(500, gin.H{
			"error": http.StatusText(resp.StatusCode),
		})
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error reading id service response", err)
		ctx.JSON(500, gin.H{
			"error": http.StatusText(500),
		})
		return
	}

	var idServiceResponse MessageFromIdService
	err = json.Unmarshal(body, &idServiceResponse)
	if err != nil {
		log.Println("Error unmarshalling id service response", err)
		ctx.JSON(500, gin.H{
			"error": http.StatusText(500),
		})
		return
	}

	_, err = client.Exec(ctx, "INSERT INTO links (short_url, original_url) VALUES ($1, $2)", idServiceResponse.ID, linkQuery)

	if err != nil {
		log.Println("Error inserting link into database", err)
		ctx.JSON(500, gin.H{
			"error": http.StatusText(500),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"short_url": fmt.Sprintf("http://localhost:3005/%s", idServiceResponse.ID), // TODO: get this from env
	})
}
