package links

import (
	"encoding/json"
	"fmt"
	"io"
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
		ctx.JSON(400, gin.H{
			"error": "link is required",
		})
		return
	}

	if !isValidHTTPSUrl(linkQuery) {
		ctx.JSON(400, gin.H{
			"error": "link must be a valid HTTPS URL",
		})
		return
	}

	client := db.GetConnection()

	err := client.Ping(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "database connection error",
		})
		return
	}

	idServiceUrl := os.Getenv("ID_SERVICE_URL")
	if idServiceUrl == "" {
		ctx.JSON(500, gin.H{
			"error": "id service url is not set",
		})
		return
	}
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/id?link=%s", idServiceUrl, linkQuery))
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "id service error",
		})
		return
	}

	if resp.StatusCode != 200 {
		ctx.JSON(500, gin.H{
			"error": "id service error",
		})
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "id service error",
		})
		return
	}

	var idServiceResponse MessageFromIdService
	err = json.Unmarshal(body, &idServiceResponse)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "id service error",
		})
		return
	}

	_, err = client.Exec(ctx, "INSERT INTO links (short_url, original_url) VALUES ($1, $2)", idServiceResponse.ID, linkQuery)

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "database error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"short_url": fmt.Sprintf("http://localhost:3000/%s", idServiceResponse.ID),
	})
}
