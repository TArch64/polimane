package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) HasAccess(ctx context.Context, userID, schemaID model.ID, access model.AccessLevel) error {
	var exists bool

	err := c.db.
		WithContext(ctx).
		Model(&model.UserSchema{}).
		Select("1 AS exists").
		Where("user_id = ? AND schema_id = ? AND access >= ?", userID, schemaID, access).
		Pluck("exists", &exists).
		Error

	if err != nil {
		return err
	}

	if !exists {
		return gorm.ErrRecordNotFound
	}

	return nil
}
