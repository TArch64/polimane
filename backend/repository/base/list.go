package repositorybase

import (
	"context"

	"polimane/backend/repository"

	"gorm.io/gorm"
)

func (c *Client[M]) List(ctx context.Context, scopes ...repository.Scope) ([]*M, error) {
	var out []*M
	err := c.ListOut(ctx, &out, scopes...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client[M]) ListOut(ctx context.Context, out interface{}, scopes ...repository.Scope) error {
	err := gorm.
		G[*M](c.DB).
		Scopes(scopes...).
		Scan(ctx, out)

	if err != nil {
		return err
	}

	return repository.DoAfterScan(out)
}

func (c *Client[M]) ListOutTx(ctx context.Context, tx *gorm.DB, out interface{}, scopes ...repository.Scope) error {
	err := gorm.
		G[*M](tx).
		Scopes(scopes...).
		Scan(ctx, out)

	if err != nil {
		return err
	}

	return repository.DoAfterScan(out)
}
