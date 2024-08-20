package service

import (
	"github.com/amiranbari/challenge/config"
	"github.com/amiranbari/challenge/repository/mongodb"
	userservice "github.com/amiranbari/challenge/service/user"
	uservalidator "github.com/amiranbari/challenge/validator/user"
)

type Service struct {
	UserSvc userservice.Service
}

func New(cfg config.Config, db *mongodb.DB) *Service {
	return &Service{
		UserSvc: userservice.New(db, uservalidator.New()),
	}
}
