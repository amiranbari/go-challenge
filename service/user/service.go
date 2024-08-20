package userservice

import (
	"context"
	"github.com/amiranbari/challenge/entity"
	params "github.com/amiranbari/challenge/param"
	validator "github.com/amiranbari/challenge/validator/user"
)

type Repository interface {
	GetAllUsers(ctx context.Context, filter params.FilterRequest) ([]entity.User, error)
}

type Service struct {
	repo Repository
	vld  validator.Validator
}

func New(repo Repository, vld validator.Validator) Service {
	return Service{
		repo: repo,
		vld:  vld,
	}
}
