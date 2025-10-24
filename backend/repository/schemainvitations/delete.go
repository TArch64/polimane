package schemainvitations

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type DeleteOptions struct {
	Email    string
	SchemaID model.ID
}

func (c *Client) Delete(ctx context.Context, options *DeleteOptions) error {
	_, err := gorm.
		G[model.SchemaInvitation](c.db).
		Where("email = ? AND schema_id = ?", options.Email, options.SchemaID).
		Delete(ctx)

	return err
}
