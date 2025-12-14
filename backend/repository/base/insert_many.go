package repositorybase

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/repository"
)

func (c *Client[M]) InsertMany(ctx context.Context, data *[]M, clauses ...clause.Expression) error {
	return c.InsertManyTx(ctx, c.DB, data, clauses...)
}

func (c *Client[M]) InsertManyTx(ctx context.Context, tx *gorm.DB, data *[]M, clauses ...clause.Expression) error {
	return gorm.
		G[M](tx, clauses...).
		CreateInBatches(ctx, data, repository.DefaultBatch)
}
