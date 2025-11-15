package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) CreateTx(ctx context.Context, tx *gorm.DB, userSchema *model.UserSchema) error {
	return gorm.
		G[model.UserSchema](tx, gorm.WithResult()).
		Create(ctx, userSchema)
}
