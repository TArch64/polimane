package repositorybase

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/repository"
)

func (c *Client[M]) Update(ctx context.Context, updates M, scopes ...repository.Scope) error {
	return c.UpdateTx(ctx, c.DB, updates, scopes...)
}

func (c *Client[M]) UpdateTx(ctx context.Context, tx *gorm.DB, updates M, scopes ...repository.Scope) error {
	_, err := gorm.
		G[M](tx).
		Scopes(scopes...).
		Updates(ctx, updates)

	return err
}
