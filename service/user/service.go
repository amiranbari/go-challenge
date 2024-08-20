package userservice

import (
	"context"
	"github.com/amiranbari/challenge/entity"
	params "github.com/amiranbari/challenge/param"
)

type Repository interface {
	GetAllUsers(ctx context.Context, filter params.FilterRequest) ([]entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
