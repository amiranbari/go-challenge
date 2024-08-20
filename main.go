package main

import (
	"github.com/amiranbari/challenge/config"
	httpserver "github.com/amiranbari/challenge/delivery/http_server"
	"github.com/amiranbari/challenge/repository/mongodb"
	"github.com/amiranbari/challenge/service"
)

func main() {
	cfg := Config()
	db := mongodb.New(cfg.MongoDB)
	svc := Service(cfg, db)
	httpServer := HTTPServer(cfg, svc)
	httpServer.Serve()
}

func Config() config.Config {
	return config.C()
}

func Service(cfg config.Config, db *mongodb.DB) *service.Service {
	return service.New(cfg, db)
}

func HTTPServer(cfg config.Config, svc *service.Service) *httpserver.Server {
	return httpserver.New(cfg, svc)
}
