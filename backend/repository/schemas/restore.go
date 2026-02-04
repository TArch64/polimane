package schemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) Restore(ctx context.Context, IDs []model.ID) error {
	return c.DB.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			return c.RestoreTx(ctx, tx, IDs)
		})
}

func (c *Client) RestoreTx(ctx context.Context, tx *gorm.DB, IDs []model.ID) error {
	err := c.Client.RestoreTx(ctx, tx,
		repository.IDsIn(IDs),
	)
	if err != nil {
		return err
	}

	return c.userSchemas.RestoreTx(ctx, tx,
		repository.SchemaIDsIn(IDs),
	)
}
