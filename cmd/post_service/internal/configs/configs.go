package configs

import (
	"iman/pkg/configer"
	"path/filepath"
	"runtime"
)

type Configs interface {
	Peek() *configs
}

const configFileName = "configs.json"

var ConfigPath = func() string {
	_, p, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	return filepath.Dir(p) + "/" + configFileName
}()

func New() Configs {
	cfg := &configs{}
	err := configer.LoadConfig(ConfigPath, cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func (c *configs) Peek() *configs {
	return c
}

type configs struct {
	Database    Database    `json:"database"`
	PostService PostService `json:"postService"`
}

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type PostService struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
