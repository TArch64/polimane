package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) DeleteSoftTx(ctx context.Context, tx *gorm.DB, scopes ...repository.Scope) error {
	scopes = append(scopes, repository.IncludeSoftDeleted)

	return c.UpdateTx(ctx, tx,
		model.UserSchema{
			SoftDeletable: model.SoftDeletedNow(),
			FolderID:      model.NilID,
		},
		scopes...,
	)
}
