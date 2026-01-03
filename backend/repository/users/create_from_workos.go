package users

import (
	"context"
	"time"

	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) CreateFromWorkos(ctx context.Context, workosUser *usermanagement.User) (user *model.User, err error) {
	user = &model.User{
		WorkosID:  workosUser.ID,
		Email:     workosUser.Email,
		FirstName: workosUser.FirstName,
		LastName:  workosUser.LastName,
	}

	err = c.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err = c.InsertTx(ctx, tx, user); err != nil {
			return err
		}

		if err = c.createSubscription(ctx, tx, user); err != nil {
			return err
		}

		return c.acceptInvitations(ctx, tx, user)
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *Client) createSubscription(ctx context.Context, tx *gorm.DB, user *model.User) error {
	return c.userSubscriptions.InsertTx(ctx, tx, &model.UserSubscription{
		ID:             user.ID,
		Plan:           model.SubscriptionBasic,
		TrialStartedAt: time.Now(),
		TrialEndsAt:    time.Now().Add(model.SubscriptionTrialDuration),
	})
}

func (c *Client) acceptInvitations(ctx context.Context, tx *gorm.DB, user *model.User) error {
	schemaInvitations, err := c.schemaInvitations.List(ctx,
		repository.EmailEq(user.Email),
	)
	if err != nil {
		return err
	}
	if len(schemaInvitations) == 0 {
		return nil
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

	return c.schemaInvitations.DeleteTx(ctx, tx,
		repository.EmailEq(user.Email),
	)
}
