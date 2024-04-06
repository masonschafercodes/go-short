package links

type Link struct {
	ID          int    `json:"id"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	CreatedAt   string `json:"created_at"`
	AccessCount int    `json:"access_count"`
}
