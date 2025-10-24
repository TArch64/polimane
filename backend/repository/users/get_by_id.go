package users

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) GetByID(ctx context.Context, id model.ID) (*model.User, error) {
	return gorm.
		G[*model.User](c.db).
		Where("id = ?", id).
		Take(ctx)
}
