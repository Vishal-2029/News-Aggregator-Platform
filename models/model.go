package models

type NewsItem struct {
	ArticleId   string   `json:"article_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	ImageURl    string   `json:"image_url"`
	PublishedAt string   `json:"pubDate"`
	Creator     []string `json:"creator"`
	Source      string   `json:"source_id"`
}

type NewsResponse struct {
	Status       string     `json:"status"`
	TotalResults int        `json:"totalResults"`
	Results      []NewsItem `json:"results"`
}
