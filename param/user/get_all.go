package userparam

import (
	"github.com/amiranbari/challenge/entity"
	"github.com/amiranbari/challenge/param"
)

type GetAllRequest struct {
	Filter param.FilterRequest
}

type GetAllResponse struct {
	Link        string
	Users       []entity.User
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
