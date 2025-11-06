package schemainvitations

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) DeleteManyByEmailTx(ctx context.Context, tx *gorm.DB, email string) error {
	_, err := gorm.
		G[model.SchemaInvitation](tx).
		Where("email = ?", email).
		Delete(ctx)

	return err
}
