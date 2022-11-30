package main

import (
	"context"
	"database/sql"
	"fmt"
	"iman/internal/services/crawler"
	"iman/pkg/proto/crawler_service"
	"log"
	"net"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"database", "5432", "postgres", "postgres", "iman")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	c := crawler.New(db)

	grpcServer := grpc.NewServer()

	crawler_service.RegisterCrawlerServiceServer(grpcServer, c)

	err = c.PopulateData(context.Background())
	if err != nil {
		fmt.Println("error populating data:", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
