package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type CreateOptions struct {
	UserID   model.ID
	SchemaID model.ID
	Access   model.AccessLevel
}

func (c *Client) CreateTx(ctx context.Context, tx *gorm.DB, userSchema *model.UserSchema) error {
	result := gorm.WithResult()

	return gorm.
		G[model.UserSchema](tx, result).
		Create(ctx, userSchema)
}
