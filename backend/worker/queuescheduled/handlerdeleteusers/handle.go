package handlerdeleteusers

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
	"polimane/backend/worker/events"
)

func (h *Handler) Handle(ctx context.Context, _ *events.Message) error {
	users, err := h.getDeletedUsers(ctx)
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return nil
	}

	for _, user := range users {
		err = h.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			return h.deleteUser(ctx, tx, user)
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) getDeletedUsers(ctx context.Context) ([]*model.User, error) {
	return h.users.List(ctx,
		repository.SoftDeleted(30),
	)
}
