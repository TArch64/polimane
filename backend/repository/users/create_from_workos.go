package users

import (
	"context"

	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) CreateFromWorkos(ctx context.Context, workosUser *usermanagement.User) (*model.User, error) {
	schemaInvitations, err := c.schemaInvitations.List(ctx, repository.EqEmail(workosUser.Email))
	if err != nil {
		return nil, err
	}

	user := &model.User{
		WorkosID:  workosUser.ID,
		Email:     workosUser.Email,
		FirstName: workosUser.FirstName,
		LastName:  workosUser.LastName,
	}

	if len(schemaInvitations) == 0 {
		err = c.createTx(ctx, c.db, user)
	} else {
		err = c.db.Transaction(func(tx *gorm.DB) error {
			if err = c.createTx(ctx, tx, user); err != nil {
				return err
			}

			userSchemas := make([]model.UserSchema, len(schemaInvitations))

			for i, invitation := range schemaInvitations {
				userSchemas[i] = model.UserSchema{
					UserID:   user.ID,
					SchemaID: invitation.SchemaID,
					Access:   invitation.Access,
				}
			}

			err = c.userSchemas.CreateManyTx(ctx, tx, &userSchemas)
			if err != nil {
				return err
			}

			return c.schemaInvitations.DeleteManyTx(ctx, tx, repository.EqEmail(workosUser.Email))
		})
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *Client) createTx(ctx context.Context, tx *gorm.DB, user *model.User) error {
	result := gorm.WithResult()

	return gorm.
		G[model.User](tx, result).
		Create(ctx, user)
}
