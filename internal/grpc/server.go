package grpc

import (
	"context"
	"database/sql"

	"github.com/Vishal-2029/news-service/pb"
)

type Server struct {
	pb.UnimplementedNewsServiceServer
	DB *sql.DB
}

func (s *Server) GetNews(ctx context.Context, req *pb.NewsRequest) (*pb.NewsResponse, error){
	rows, err := s.DB.Query("SELECT articleid, title, description, link, image_url ,published_at, source FROM news")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*pb.NewsItem
	for rows.Next() {
		var n pb.NewsItem
		err := rows.Scan(&n.ArticleId, &n.Title, &n.Description, &n.Link, &n.ImageURl, &n.PublishedAt, &n.Source)
		if err == nil {
			items = append(items, &n)
		}
	}

	return &pb.NewsResponse{Items: items}, nil
}