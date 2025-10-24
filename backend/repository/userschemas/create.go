package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type CreateOptions struct {
	UserID   model.ID
	SchemaID model.ID
	Access   model.AccessLevel
}

func (c *Client) CreateTx(ctx context.Context, tx *gorm.DB, options *CreateOptions) (*model.UserSchema, error) {
	userSchema := &model.UserSchema{
		UserID:   options.UserID,
		SchemaID: options.SchemaID,
		Access:   options.Access,
	}

	result := gorm.WithResult()

	err := gorm.
		G[model.UserSchema](tx, result).
		Create(ctx, userSchema)

	if err != nil {
		return nil, err
	}

	return userSchema, nil
}

type CreateWithAccessCheckOptions = WithAccessCheck[CreateOptions]

func (c *Client) CreateWithAccessCheck(ctx context.Context, options *CreateWithAccessCheckOptions) (*model.UserSchema, error) {
	err := c.HasAccess(ctx, options.CurrentUser.ID, options.Operation.SchemaID, model.AccessAdmin)
	if err != nil {
		return nil, err
	}

	return c.CreateTx(ctx, c.db, options.Operation)
}

func (c *Client) CreateManyTx(ctx context.Context, tx *gorm.DB, items []*CreateOptions) error {
	userSchemas := make([]model.UserSchema, len(items))

	for i, item := range items {
		userSchemas[i] = model.UserSchema{
			UserID:   item.UserID,
			SchemaID: item.SchemaID,
			Access:   item.Access,
		}
	}

	return gorm.
		G[model.UserSchema](tx).
		CreateInBatches(ctx, &userSchemas, 100)
}
