package repositorybase

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/repository"
)

func (c *Client[M]) Get(ctx context.Context, scopes ...repository.Scope) (*M, error) {
	return gorm.
		G[*M](c.DB).
		Scopes(scopes...).
		Take(ctx)
}

func (c *Client[M]) GetOut(ctx context.Context, out interface{}, scopes ...repository.Scope) error {
	return gorm.
		G[*M](c.DB).
		Scopes(scopes...).
		Scan(ctx, out)
}
