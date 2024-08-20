package mongodb

import (
	"context"
	"github.com/amiranbari/challenge/entity"
	params "github.com/amiranbari/challenge/param"
)

func (d *DB) GetAllUsers(ctx context.Context, filter params.FilterRequest) ([]entity.User, error) {
	users := make([]entity.User, 0)
	return users, nil
}
