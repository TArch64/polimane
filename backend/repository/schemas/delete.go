package schemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type DeleteOptions struct {
	Ctx      context.Context
	User     *model.User
	SchemaID model.ID
}

func (c *Impl) Delete(options *DeleteOptions) (err error) {
	err = c.userSchemas.HasAccess(options.Ctx, options.User.ID, options.SchemaID)
	if err != nil {
		return err
	}

	err = c.db.WithContext(options.Ctx).Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(&model.Schema{}, options.SchemaID).Error; err != nil {
			return err
		}

		return c.userSchemas.DeleteTx(tx, options.User.ID, options.SchemaID)
	})

	if err != nil {
		return err
	}

	c.signals.InvalidateUserCache.Emit(options.Ctx, options.User.ID)
	return nil
}
