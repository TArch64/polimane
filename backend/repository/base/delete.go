package repositorybase

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/repository"
)

func (c *Client[M]) Delete(ctx context.Context, scopes ...repository.Scope) error {
	return c.DeleteTx(ctx, c.DB, scopes...)
}

func (c *Client[M]) DeleteTx(ctx context.Context, tx *gorm.DB, scopes ...repository.Scope) error {
	affected, err := gorm.
		G[M](tx).
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
