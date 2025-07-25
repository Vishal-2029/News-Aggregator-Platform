package db

import (
	"database/sql"
	"log"

	"github.com/Vishal-2029/models"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB(dsn string) error {
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return db.Ping()
}

func GetDB() *sql.DB {
	return db
}


func GetAllNews() ([]models.NewsItem, error) {
	rows, err := db.Query(`
		SELECT Articleid, title, description, link, image_url, published_at, source
		FROM news
		ORDER BY published_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.NewsItem
	for rows.Next() {
		var n models.NewsItem
		if err := rows.Scan(
			&n.ArticleId,
			&n.Title,
			&n.Description,
			&n.Link,
			&n.ImageURl,
			&n.PublishedAt,
			&n.Source,
		); err != nil {
			return nil, err
		}
		list = append(list, n)
	}
	return list, rows.Err()
}

func SaveNews(news models.NewsItem) error {
	query := `INSERT INTO news (Articleid, title, description, link, image_url, published_at, source)
	          VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(query,
		news.ArticleId,
		news.Title,
		news.Description,
		news.Link,
		news.ImageURl,
		news.PublishedAt,
		news.Source,
	)
	if err != nil {
		log.Println("Failed to insert into DB:", err)
	}
	return err
}
