package server

import (
	"github.com/developerkaren96/MetricAlertPro/internal/storage"
)

type Config struct {
	ListenAddress string
}

func NewServerConfig() Config {
	return Config{
		ListenAddress: "0.0.0.0:8080",
	}
}

type Server struct {
	Config  Config
	Storage storage.Storage
}

func NewServer() *Server {
	return &Server{
		Config:  NewServerConfig(),
		Storage: storage.NewMemStorage(),
	}
}
