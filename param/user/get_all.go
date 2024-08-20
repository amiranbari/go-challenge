package userparam

import (
	"github.com/amiranbari/challenge/param"
)

type GetAllRequest struct {
	Filter param.FilterRequest
}

type GetAllResponse struct {
	Link string
}
