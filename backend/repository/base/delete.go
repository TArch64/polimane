package repositorybase

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/repository"
)

func (c *Client[M]) Delete(ctx context.Context, scopes ...repository.Scope) error {
	return c.DeleteTx(ctx, c.DB, scopes...)
}

func (c *Client[M]) DeleteCounted(ctx context.Context, scopes ...repository.Scope) (int, error) {
	return c.DeleteCountedTx(ctx, c.DB, scopes...)
}

func (c *Client[M]) DeleteTx(ctx context.Context, tx *gorm.DB, scopes ...repository.Scope) error {
	_, err := c.DeleteCountedTx(ctx, tx, scopes...)
	return err
}

func (c *Client[M]) DeleteCountedTx(ctx context.Context, tx *gorm.DB, scopes ...repository.Scope) (int, error) {
	affected, err := gorm.
		G[M](tx).
		Scopes(scopes...).
		Delete(ctx)

	if err != nil {
		return 0, err
	}

	if affected == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	return affected, nil
}
