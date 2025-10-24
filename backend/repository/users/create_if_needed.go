package users

import (
	"context"
	"errors"

	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) CreateIfNeeded(ctx context.Context, workosUser *usermanagement.User) (*model.User, error) {
	user, err := c.GetByWorkosID(ctx, workosUser.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.CreateFromWorkos(ctx, workosUser)
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
