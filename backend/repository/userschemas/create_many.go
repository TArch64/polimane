package userschemas

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/model"
)

func (c *Client) CreateMany(ctx context.Context, userSchemas *[]model.UserSchema) error {
	return c.CreateManyTx(ctx, c.db, userSchemas)
}

func (c *Client) CreateManyTx(ctx context.Context, tx *gorm.DB, userSchemas *[]model.UserSchema) error {
	return gorm.
		G[model.UserSchema](tx, clause.OnConflict{DoNothing: true}).
		CreateInBatches(ctx, userSchemas, 100)
}
