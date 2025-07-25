package fetcher

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Vishal-2029/models"
	"github.com/nats-io/nats.go"
)

func FetchNews(nc *nats.Conn) []models.NewsItem {
	resp, err := http.Get("https://newsdata.io/api/1/news?apikey=pub_38a9bb63b9e547878d4e1cb207a6be5f&q=cryptocurrency%20regulations&language=en")
	if err != nil {
		log.Println("Error fetching news:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Failed to fetch news. Status code:", resp.StatusCode)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil
	}

	var news models.NewsResponse
	if err := json.Unmarshal(body, &news); err != nil {
		log.Println("Error decoding JSON:", err)
		return nil
	}

	return news.Results
}
