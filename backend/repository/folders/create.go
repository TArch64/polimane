package folders

import (
	"context"

	"polimane/backend/model"

	"gorm.io/gorm"
)

func (c *Client) CreateTx(ctx context.Context, tx *gorm.DB, folder *model.Folder) error {
	return gorm.
		G[model.Folder](tx, gorm.WithResult()).
		Create(ctx, folder)
}
