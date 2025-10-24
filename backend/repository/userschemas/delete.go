package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type DeleteOptions struct {
	UserID   model.ID
	SchemaID model.ID
}

func (c *Client) DeleteTx(ctx context.Context, tx *gorm.DB, options *DeleteOptions) error {
	_, err := gorm.
		G[model.UserSchema](tx).
		Where("user_id = ? AND schema_id = ?", options.UserID, options.SchemaID).
		Delete(ctx)

	return err
}

type DeleteWithAccessCheckOptions = WithAccessCheck[DeleteOptions]

func (c *Client) DeleteWithAccessCheck(ctx context.Context, options *DeleteWithAccessCheckOptions) error {
	err := c.HasAccess(ctx, options.CurrentUser.ID, options.Operation.SchemaID, model.AccessAdmin)
	if err != nil {
		return nil
	}

	return c.DeleteTx(ctx, c.db, options.Operation)
}
