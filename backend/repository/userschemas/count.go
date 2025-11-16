package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) Count(ctx context.Context, scopes ...repository.Scope) (int64, error) {
	return gorm.
		G[model.UserSchema](c.db).
		Scopes(scopes...).
		Count(ctx, "schema_id")
}
