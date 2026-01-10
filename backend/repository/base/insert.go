package repositorybase

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c *Client[M]) Insert(ctx context.Context, data *M, clauses ...clause.Expression) error {
	return c.InsertTx(ctx, c.DB, data, clauses...)
}

func (c *Client[M]) InsertTx(ctx context.Context, tx *gorm.DB, data *M, clauses ...clause.Expression) error {
	return gorm.
		G[M](tx, append(clauses, gorm.WithResult())...).
		Create(ctx, data)
}
