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
	err := c.ListOut(ctx, schemaIDs,
		repository.IncludeSoftDeleted, // required for deleted schemas
		repository.Select("schema_id"),
		repository.UserIDEq(user.ID),
		repository.SchemaIDsIn(*schemaIDs),
		repository.Where("access >= ?", access),
	)

	if err != nil {
		return err
	}

	if len(*schemaIDs) == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
