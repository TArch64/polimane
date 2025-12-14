package users

import (
	"context"

	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) CreateFromWorkos(ctx context.Context, workosUser *usermanagement.User) (*model.User, error) {
	schemaInvitations, err := c.schemaInvitations.List(ctx,
		repository.EmailEq(workosUser.Email),
	)
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
		err = c.Insert(ctx, user)
	} else {
		err = c.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			if err = c.InsertTx(ctx, tx, user); err != nil {
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

			err = c.userSchemas.InsertManyTx(ctx, tx, &userSchemas,
				clause.OnConflict{DoNothing: true},
			)
			if err != nil {
				return err
			}

			return c.schemaInvitations.DeleteTx(ctx, tx, repository.EmailEq(workosUser.Email))
		})
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
