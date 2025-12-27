package repositorybase

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/repository"
)

func (c *Client[M]) Restore(ctx context.Context, scopes ...repository.Scope) error {
	return c.RestoreTx(ctx, c.DB, scopes...)
}

func (c *Client[M]) RestoreTx(ctx context.Context, tx *gorm.DB, scopes ...repository.Scope) error {
	scopes = append(scopes, repository.IncludeSoftDeleted)

	affected, err := gorm.
		G[M](tx).
		Scopes(scopes...).
		Update(ctx, "deleted_at", nil)

	if err != nil {
		return err
	}
	if affected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
