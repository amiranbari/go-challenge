package config

import "github.com/amiranbari/challenge/repository/mongodb"

type HTTPServer struct {
	Port int `koanf:"port"`
}

type Config struct {
	HTTPServer HTTPServer     `koanf:"http_server"`
	MongoDB    mongodb.Config `koanf:"mongodb"`
}
