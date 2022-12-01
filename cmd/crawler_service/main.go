package main

import (
	"context"
	"database/sql"
	"fmt"
	"iman/cmd/crawler_service/internal/configs"
	"iman/internal/services/crawler"
	"iman/pkg/proto/crawler_service"
	"log"
	"net"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

func main() {
	cfg := configs.New()

	fmt.Println(cfg)

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

	c := crawler.New(db)

	grpcServer := grpc.NewServer()

	crawler_service.RegisterCrawlerServiceServer(grpcServer, c)

	err = c.PopulateData(context.Background())
	if err != nil {
		log.Fatalln("error populating data:", err)
	}

	lis, err := net.Listen("tcp", cfg.CrawlerService.Port)
	if err != nil {
		log.Fatalf("Failed to listen on port %v: %v", err, cfg.CrawlerService.Port)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %v: %v", err, cfg.CrawlerService.Port)
	}
}
