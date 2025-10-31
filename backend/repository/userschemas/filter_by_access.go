package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) FilterByAccess(
	ctx context.Context,
	userID model.ID,
	schemaIDs []model.ID,
	access model.AccessLevel,
) ([]model.ID, error) {
	var allowedIDs []model.ID

	err := gorm.
		G[model.UserSchema](c.db).
		Select("schema_id").
		Where("user_id = ? AND schema_id IN (?) AND access >= ?", userID, schemaIDs, access).
		Scan(ctx, &allowedIDs)

	return allowedIDs, err
}
