package main

import (
	"iman/cmd/api_gateway/internal/handlers"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	conn := postServiceConnection()
	defer conn.Close()

	h := handlers.NewHandler(conn)

	r.GET("/posts", h.GetPosts)
	r.GET("/post", h.GetPostByID)
	r.DELETE("/post", h.DeletePostByID)
	r.PUT("/post", h.UpdatePostByID)

	r.Run(":8080")
}

func postServiceConnection() *grpc.ClientConn {
	conn, err := grpc.Dial("post_service:9001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return conn
}
