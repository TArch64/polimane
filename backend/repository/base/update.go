package repositorybase

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/repository"
)

func (c *Client[M]) Update(ctx context.Context, updates M, scopes ...repository.Scope) error {
	return c.UpdateTx(ctx, c.DB, updates, scopes...)
}

func (c *Client[M]) UpdateColumn(ctx context.Context, column string, value any, scopes ...repository.Scope) error {
	return c.UpdateColumnTx(ctx, c.DB, column, value, scopes...)
}

func (c *Client[M]) UpdateTx(ctx context.Context, tx *gorm.DB, updates M, scopes ...repository.Scope) (err error) {
	_, err = gorm.
		G[M](tx).
		Scopes(scopes...).
		Updates(ctx, updates)

	return
}

func (c *Client[M]) UpdateColumnTx(ctx context.Context, tx *gorm.DB, column string, value any, scopes ...repository.Scope) (err error) {
	_, err = gorm.
		G[M](tx).
		Scopes(scopes...).
		Update(ctx, column, value)

	return
}
