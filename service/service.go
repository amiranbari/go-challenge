package service

import (
	"github.com/amiranbari/challenge/config"
	"github.com/amiranbari/challenge/repository/mongodb"
	userservice "github.com/amiranbari/challenge/service/user"
)

type Service struct {
	UserSvc userservice.Service
}

func New(cfg config.Config, db *mongodb.DB) *Service {
	return &Service{
		UserSvc: userservice.New(db),
	}
}
