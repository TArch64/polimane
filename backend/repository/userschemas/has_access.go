package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) HasAccess(
	ctx context.Context,
	userID, schemaID model.ID,
	access model.AccessLevel,
) error {
	exists, err := c.Exists(ctx,
		repository.Where("user_id = ? AND schema_id = ? AND access >= ?", userID, schemaID, access),
	)

	if err != nil {
		return err
	}

	if !exists {
		return gorm.ErrRecordNotFound
	}

	return nil
}
