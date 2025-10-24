package users

import (
	"context"

	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/model"
)

func (c *Client) CreateIfNeeded(ctx context.Context, workosUser usermanagement.User) (*model.User, error) {
	user := model.User{
		WorkosID:  workosUser.ID,
		Email:     workosUser.Email,
		FirstName: workosUser.FirstName,
		LastName:  workosUser.LastName,
	}

	err := c.db.
		WithContext(ctx).
		Where(user).
		Attrs(user).
		FirstOrCreate(&user).
		Error

	return &user, err
}
