package main

import (
	"database/sql"
	"fmt"
	"iman/cmd/post_service/internal/configs"
	"iman/internal/services/post"
	pb "iman/pkg/proto/post_service"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	cfg := configs.New()

	lis, err := net.Listen("tcp", cfg.PostService.Port)
	if err != nil {
		log.Fatalf("Failed to listen on port %v: %v", err, cfg.PostService.Port)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	s := post.New(db)

	grpcServer := grpc.NewServer()

	pb.RegisterPostServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %v: %v", err, cfg.PostService.Port)
	}
}
