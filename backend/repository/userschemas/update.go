package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type UpdateOptions struct {
	UserID   model.ID
	SchemaID model.ID
	Updates  *model.UserSchema
}

func (c *Client) Update(ctx context.Context, options *UpdateOptions) error {
	_, err := gorm.
		G[model.UserSchema](c.db).
		Where("user_id = ? AND schema_id = ?", options.UserID, options.SchemaID).
		Updates(ctx, *options.Updates)

	return err
}

type UpdateWithAccessCheckOptions = WithAccessCheck[UpdateOptions]

func (c *Client) UpdateWithAccessCheck(ctx context.Context, options *UpdateWithAccessCheckOptions) error {
	err := c.HasAccess(ctx, options.CurrentUser.ID, options.Operation.SchemaID, model.AccessAdmin)
	if err != nil {
		return nil
	}

	return c.Update(ctx, options.Operation)
}
