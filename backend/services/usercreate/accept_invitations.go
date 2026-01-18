package usercreate

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/model"
	"polimane/backend/repository"
	"polimane/backend/services/subscriptioncounters"
)

func (s *Service) acceptInvitations(ctx context.Context, tx *gorm.DB, user *model.User) error {
	schemaInvitations, err := s.schemaInvitations.List(ctx,
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

	insertResult := gorm.WithResult()
	err = s.userSchemas.InsertManyTx(ctx, tx, &userSchemas,
		clause.OnConflict{DoNothing: true},
		insertResult,
	)
	if err != nil {
		return err
	}

	err = s.schemaInvitations.DeleteTx(ctx, tx,
		repository.EmailEq(user.Email),
	)
	if err != nil {
		return err
	}

	return s.subscriptionCounters.SchemasCreated.ChangeTx(ctx, tx, subscriptioncounters.ChangeSet{
		user.ID: int16(insertResult.RowsAffected),
	})
}
