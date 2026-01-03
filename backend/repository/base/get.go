package repositorybase

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/repository"
)

func (c *Client[M]) Get(ctx context.Context, scopes ...repository.Scope) (*M, error) {
	return c.GetCustomizable(ctx, nil, scopes...)
}

func (c *Client[M]) GetCustomizable(ctx context.Context, customizer Customizer[*M], scopes ...repository.Scope) (*M, error) {
	chain := gorm.G[*M](c.DB).Scopes(scopes...)
	if customizer != nil {
		chain = customizer(chain)
	}
	return chain.Take(ctx)
}

func (c *Client[M]) GetOut(ctx context.Context, out interface{}, scopes ...repository.Scope) error {
	return gorm.
		G[*M](c.DB).
		Scopes(scopes...).
		Scan(ctx, out)
}
