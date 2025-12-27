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
			err := c.RestoreTx(ctx, tx,
				repository.IDsIn(IDs),
			)
			if err != nil {
				return err
			}

			return c.userSchemas.RestoreTx(ctx, tx,
				repository.SchemaIDsIn(IDs),
			)
		})
}
