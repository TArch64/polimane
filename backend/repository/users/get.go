package users

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) Get(ctx context.Context, scopes ...repository.Scope) (*model.User, error) {
	return gorm.
		G[*model.User](c.DB).
		Scopes(scopes...).
		Take(ctx)
}
