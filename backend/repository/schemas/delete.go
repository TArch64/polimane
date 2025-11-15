package schemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) Delete(ctx context.Context, scopes ...repository.Scope) error {
	_, err := gorm.
		G[model.Schema](c.db).
		Scopes(scopes...).
		Delete(ctx)

	return err
}
