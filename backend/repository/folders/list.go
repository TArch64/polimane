package folders

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/repository"

	"gorm.io/gorm"
)

func (c *Client) List(ctx context.Context, scopes ...repository.Scope) ([]*model.Folder, error) {
	return gorm.
		G[*model.Folder](c.db).
		Scopes(scopes...).
		Find(ctx)
}
