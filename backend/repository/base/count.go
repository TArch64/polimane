package repositorybase

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/repository"
)

func (c *Client[M]) Count(ctx context.Context, scopes ...repository.Scope) (int64, error) {
	return c.CountByColumn(ctx, "id", scopes...)
}

func (c *Client[M]) CountByColumn(ctx context.Context, column string, scopes ...repository.Scope) (int64, error) {
	return gorm.
		G[M](c.DB).
		Scopes(scopes...).
		Count(ctx, column)
}
