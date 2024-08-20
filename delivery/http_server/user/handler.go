package userhandler

import userservice "github.com/amiranbari/challenge/service/user"

type Handler struct {
	userSvc userservice.Service
}

func New(
	userSvc userservice.Service,
) Handler {
	return Handler{
		userSvc: userSvc,
	}
}
