package schemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) Count(ctx context.Context, scopes ...repository.Scope) (int64, error) {
	return gorm.
		G[model.Schema](c.db).
		Scopes(scopes...).
		Count(ctx, "id")
}
