package users

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) GeyByEmail(ctx context.Context, email string) (*model.User, error) {
	return gorm.
		G[*model.User](c.db).
		Where("email = ?", email).
		Take(ctx)
}
