package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) FilterByAccess(
	ctx context.Context,
	user *model.User,
	schemaIDs *[]model.ID,
	access model.AccessLevel,
) error {
	return gorm.
		G[model.UserSchema](c.db).
		Select("schema_id").
		Where("user_id = ? AND schema_id IN (?) AND access >= ?", user.ID, *schemaIDs, access).
		Scan(ctx, schemaIDs)
}
