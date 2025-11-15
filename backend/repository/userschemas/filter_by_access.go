package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) FilterByAccess(
	ctx context.Context,
	user *model.User,
	schemaIDs *[]model.ID,
	access model.AccessLevel,
) error {
	err := gorm.
		G[model.UserSchema](c.db).
		Select("schema_id").
		Scopes(
			repository.UserIDEq(user.ID),
			repository.SchemaIDsIn(*schemaIDs),
		).
		Where("access >= ?", access).
		Scan(ctx, schemaIDs)

	if err != nil {
		return err
	}

	if len(*schemaIDs) == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
