package users

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) GetByWorkosID(ctx context.Context, id string) (*model.User, error) {
	return gorm.
		G[*model.User](c.db).
		Where("workos_id = ?", id).
		Take(ctx)
}
