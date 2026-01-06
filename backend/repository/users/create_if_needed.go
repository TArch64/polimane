package users

import (
	"context"
	"errors"

	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) CreateIfNeeded(ctx context.Context, workosUser *usermanagement.User) (*model.User, *CreatingFlags, error) {
	user, err := c.Get(ctx,
		WorkosIDEq(workosUser.ID),
		repository.IncludeSoftDeleted,
	)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.CreateFromWorkos(ctx, workosUser)
	}
	if err != nil {
		return nil, nil, err
	}
	return user, &CreatingFlags{}, nil
}
