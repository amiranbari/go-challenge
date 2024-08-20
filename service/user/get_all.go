package userservice

import (
	"context"
	"errors"
	userparam "github.com/amiranbari/challenge/param/user"
)

func (s Service) GetAll(ctx context.Context, req userparam.GetAllRequest) (userparam.GetAllResponse, error) {

	_, err := s.repo.GetAllUsers(ctx, req.Filter)
	if err != nil {
		return userparam.GetAllResponse{}, errors.New("error in export excel")
	}

	return userparam.GetAllResponse{
		Link: "http://asdasd",
	}, nil
}
