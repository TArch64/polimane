package folders

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) Update(ctx context.Context, updates model.Folder, scopes ...repository.Scope) error {
	_, err := gorm.
		G[model.Folder](c.db).
		Scopes(scopes...).
		Updates(ctx, updates)

	return err
}
