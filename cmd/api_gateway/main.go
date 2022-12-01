package main

import (
	"iman/cmd/api_gateway/internal/configs"
	"iman/cmd/api_gateway/internal/handlers"
	"iman/cmd/api_gateway/internal/server"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cfg := configs.New()

	conn := postServiceConnection(cfg)
	defer conn.Close()

	h := handlers.New(handlers.Params{PostServiceClientConnection: conn})

	s := server.New(server.Params{Configs: cfg, Handlers: h})

	if err := s.Run(); err != nil {
		log.Fatalln("unable to run server:", err)
	}
}

func postServiceConnection(cfg configs.Configs) *grpc.ClientConn {
	conn, err := grpc.Dial(cfg.Peek().PostService.Host+cfg.Peek().PostService.Port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return conn
}
