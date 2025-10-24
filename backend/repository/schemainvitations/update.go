package schemainvitations

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type UpdateOptions struct {
	Email    string
	SchemaID model.ID
	Updates  *model.SchemaInvitation
}

func (c *Client) Update(ctx context.Context, options *UpdateOptions) error {
	_, err := gorm.
		G[model.SchemaInvitation](c.db).
		Where("email = ? AND schema_id = ?", options.Email, options.SchemaID).
		Updates(ctx, *options.Updates)

	return err
}
