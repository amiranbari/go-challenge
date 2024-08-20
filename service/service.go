package service

import (
	"github.com/amiranbari/challenge/config"
	"github.com/amiranbari/challenge/repository/mongodb"
)

type Service struct {
}

func New(cfg config.Config, db *mongodb.DB) *Service {
	return &Service{}
}
