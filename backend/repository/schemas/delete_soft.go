package schemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) DeleteSoft(ctx context.Context, user *model.User, IDs []model.ID) error {
	return c.DB.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			return c.DeleteSoftTx(ctx, tx, user, IDs)
		})
}

func (c *Client) DeleteSoftTx(ctx context.Context, tx *gorm.DB, user *model.User, IDs []model.ID) error {
	err := c.UpdateTx(ctx, tx,
		model.Schema{
			SoftDeletable: model.SoftDeletedNow(),
			DeletedBy:     &user.ID,
		},
		repository.IDsIn(IDs),
	)

	if err != nil {
		return err
	}

	return c.userSchemas.DeleteSoftTx(ctx, tx,
		repository.SchemaIDsIn(IDs),
	)
}
