package userschemas

import (
	"context"
	"time"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) SoftDeleteTx(ctx context.Context, tx *gorm.DB, scopes ...repository.Scope) error {
	scopes = append(scopes, repository.IncludeSoftDeleted)

	affected, err := gorm.
		G[model.UserSchema](tx).
		Scopes(scopes...).
		Updates(ctx, model.UserSchema{
			SoftDeletable: &model.SoftDeletable{
				DeletedAt: gorm.DeletedAt{
					Valid: true,
					Time:  time.Now(),
				},
			},
			FolderID: model.NilID,
		})

	if err != nil {
		return err
	}
	if affected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
