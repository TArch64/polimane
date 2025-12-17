package handlerdeleteusers

import (
	"context"
	"errors"

	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (h *Handler) deleteUser(ctx context.Context, tx *gorm.DB, user *model.User) (err error) {
	if err = h.deleteUserFolders(ctx, tx, user); err != nil {
		return err
	}

	if err = h.deleteUserSchemas(ctx, tx, user); err != nil {
		return err
	}

	if err = h.deleteWorkosUser(ctx, user); err != nil {
		return err
	}

	return h.deleteUserRecord(ctx, tx, user)
}

func (h *Handler) deleteUserFolders(ctx context.Context, tx *gorm.DB, user *model.User) error {
	return ignoreNotFound(h.folders.DeleteTx(ctx, tx,
		repository.UserIDEq(user.ID),
	))
}

func (h *Handler) deleteWorkosUser(ctx context.Context, user *model.User) error {
	return h.workos.UserManagement.DeleteUser(ctx, usermanagement.DeleteUserOpts{
		User: user.WorkosID,
	})
}

func (h *Handler) deleteUserRecord(ctx context.Context, tx *gorm.DB, user *model.User) error {
	return h.users.DeleteTx(ctx, tx,
		repository.IDEq(user.ID),
		repository.IncludeSoftDeleted,
	)
}

func ignoreNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}
