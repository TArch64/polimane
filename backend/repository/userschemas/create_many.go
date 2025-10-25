package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) CreateManyTx(ctx context.Context, tx *gorm.DB, items []*CreateOptions) error {
	userSchemas := make([]model.UserSchema, len(items))

	for i, item := range items {
		userSchemas[i] = model.UserSchema{
			UserID:   item.UserID,
			SchemaID: item.SchemaID,
			Access:   item.Access,
		}
	}

	return gorm.
		G[model.UserSchema](tx).
		CreateInBatches(ctx, &userSchemas, 100)
}
