package repositorybase

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/repository"
)

func (c *Client[M]) Get(ctx context.Context, scopes ...repository.Scope) (*M, error) {
	var out M
	err := c.GetOut(ctx, &out, scopes...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client[M]) GetOut(ctx context.Context, out interface{}, scopes ...repository.Scope) error {
	return gorm.
		G[*M](c.DB).
		Scopes(scopes...).
		Scan(ctx, out)
}
