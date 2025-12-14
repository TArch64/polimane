package folders

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) HasAccess(ctx context.Context, userID, folderID model.ID) error {
	exists, err := c.Exists(ctx,
		repository.Where("id = ? AND user_id = ?", folderID, userID),
	)

	if err != nil {
		return err
	}

	if !exists {
		return gorm.ErrRecordNotFound
	}

	return nil
}
