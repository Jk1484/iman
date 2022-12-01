package server

import (
	"iman/cmd/api_gateway/internal/configs"
	"iman/cmd/api_gateway/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run() error
}

type Params struct {
	Handlers handlers.Handlers
	Configs  configs.Configs
}

type server struct {
	handlers handlers.Handlers
	configs  configs.Configs
}

func New(p Params) Server {
	return &server{
		handlers: p.Handlers,
		configs:  p.Configs,
	}
}

func (s *server) Run() error {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/posts", s.handlers.GetPosts)
	r.GET("/post", s.handlers.GetPostByID)
	r.DELETE("/post", s.handlers.DeletePostByID)
	r.PUT("/post", s.handlers.UpdatePostByID)

	return r.Run(s.configs.Peek().Server.Port)
}
