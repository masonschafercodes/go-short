package links

import "time"

type Link struct {
	ID          int       `json:"id"`
	ShortURL    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	AccessCount int       `json:"access_count"`
}
