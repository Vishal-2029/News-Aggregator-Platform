package cmd

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/Vishal-2029/config/db"
	grpcserver "github.com/Vishal-2029/internal/grpc"
	pb "github.com/Vishal-2029/news-service/pb"
)

// StartGRPCServer starts the gRPC server on port 50051
func StartGRPCServer() {
	// Listen on port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Create gRPC server instance
	s := grpc.NewServer()

	// Register service
	pb.RegisterNewsServiceServer(s, &grpcserver.Server{DB: db.GetDB()})

	log.Println("🚀 gRPC server started on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("gRPC server failed to serve: %v", err)
	}
}
