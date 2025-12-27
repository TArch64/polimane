package schemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) SoftDelete(ctx context.Context, IDs []model.ID) error {
	return c.DB.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			err := c.DeleteTx(ctx, tx,
				repository.IDsIn(IDs),
			)
			if err != nil {
				return err
			}

			return c.userSchemas.SoftDeleteTx(ctx, tx,
				repository.SchemaIDsIn(IDs),
			)
		})
}
