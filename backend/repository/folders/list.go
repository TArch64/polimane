package folders

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) ListOut(ctx context.Context, out interface{}, scopes ...repository.Scope) (err error) {
	return gorm.
		G[model.Folder](c.db).
		Scopes(scopes...).
		Scan(ctx, out)
}
