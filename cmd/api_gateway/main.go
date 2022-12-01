package main

import (
	"iman/cmd/api_gateway/internal/configs"
	"iman/cmd/api_gateway/internal/handlers"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	cfg := configs.New()

	conn := postServiceConnection(cfg)
	defer conn.Close()

	h := handlers.NewHandler(conn)

	r.GET("/posts", h.GetPosts)
	r.GET("/post", h.GetPostByID)
	r.DELETE("/post", h.DeletePostByID)
	r.PUT("/post", h.UpdatePostByID)

	r.Run(cfg.Server.Port)
}

func postServiceConnection(cfg *configs.Configs) *grpc.ClientConn {
	conn, err := grpc.Dial(cfg.PostService.Host+cfg.PostService.Port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return conn
}
