package folders

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) HasAccess(ctx context.Context, userID, folderID model.ID) error {
	var exists bool

	err := gorm.
		G[model.Folder](c.db).
		Select("1 AS exists").
		Where("id = ? AND user_id = ?", folderID, userID).
		Scan(ctx, &exists)

	if err != nil {
		return err
	}

	if !exists {
		return gorm.ErrRecordNotFound
	}

	return nil
}
