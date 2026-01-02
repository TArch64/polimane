package handlerdeleteusers

import (
	"context"
	"errors"
	"fmt"

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

	var errs []error

	for _, user := range users {
		if err = h.deleteUser(ctx, user); err != nil {
			err = fmt.Errorf("failed to delete user %s: %w", user.ID.String(), err)
			errs = append(errs, err)
			continue
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (h *Handler) getDeletedUsers(ctx context.Context) ([]*model.User, error) {
	return h.users.List(ctx,
		repository.SoftDeletedDaysAgo(30),
	)
}

func (h *Handler) deleteUser(ctx context.Context, user *model.User) error {
	for _, deleter := range h.deleters {
		if err := deleter.Collect(ctx, user); err != nil {
			return err
		}
	}

	for _, deleter := range h.deleters {
		if err := deleter.Delete(ctx); err != nil {
			return err
		}

		deleter.LogResults(ctx)
		deleter.Cleanup()
	}

	return nil
}
