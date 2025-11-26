package folders

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) Delete(ctx context.Context, scopes ...repository.Scope) error {
	affected, err := gorm.
		G[model.Folder](c.db).
		Scopes(scopes...).
		Delete(ctx)

	if err != nil {
		return err
	}
	if affected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
