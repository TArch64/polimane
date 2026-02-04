package repositorybase

import (
	"context"
	"database/sql"

	"polimane/backend/repository"

	"gorm.io/gorm"
)

func (c *Client[M]) List(ctx context.Context, scopes ...repository.Scope) (out []*M, err error) {
	err = c.ListOut(ctx, &out, scopes...)
	return
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

func (c *Client[M]) ListRows(ctx context.Context, scopes ...repository.Scope) (rows *sql.Rows, err error) {
	rows, err = gorm.
		G[M](c.DB).
		Scopes(scopes...).
		Rows(ctx)

	return
}
