package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) DeleteTx(tx *gorm.DB, userID, schemaID model.ID) error {
	return tx.
		Where("user_id = ? AND schema_id = ?", userID, schemaID).
		Delete(&model.UserSchema{}).
		Error
}

type DeleteWithAccessCheckOptions struct {
	User     *model.User
	UserID   model.ID
	SchemaID model.ID
}

func (c *Client) DeleteWithAccessCheck(ctx context.Context, options *DeleteWithAccessCheckOptions) error {
	userSchemasQuery := gorm.G[model.UserSchema](c.db).
		Select("schema_id").
		Where("user_id = ? AND access = ?", options.User.ID, model.AccessAdmin)

	_, err := gorm.G[model.UserSchema](c.db).
		Where("schema_id = ? AND user_id = ? AND schema_id IN (?)", options.SchemaID, options.UserID, userSchemasQuery).
		Delete(ctx)

	return err
}
