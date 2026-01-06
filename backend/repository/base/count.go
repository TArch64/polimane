package repositorybase

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"polimane/backend/repository"
)

func (c *Client[M]) Count(ctx context.Context, scopes ...repository.Scope) (int64, error) {
	return c.CountByColumn(ctx, "id", scopes...)
}

func (c *Client[M]) CountQuery(ctx context.Context, scopes ...repository.Scope) gorm.ChainInterface[M] {
	return c.CountByColumnQuery(ctx, "id", scopes...)
}

func (c *Client[M]) CountByColumn(ctx context.Context, column string, scopes ...repository.Scope) (int64, error) {
	return gorm.
		G[M](c.DB).
		Scopes(scopes...).
		Count(ctx, column)
}

func (c *Client[M]) CountByColumnQuery(ctx context.Context, column string, scopes ...repository.Scope) gorm.ChainInterface[M] {
	columnSelect := fmt.Sprintf("COUNT(%s) AS count", column)

	return gorm.
		G[M](c.DB).
		Scopes(append(scopes, repository.Select(columnSelect))...)
}
