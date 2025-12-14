package handlercleanupinvitations

import (
	"context"
	"errors"

	"gorm.io/gorm"

	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	"polimane/backend/worker/events"
)

func (h *Handler) Handle(ctx context.Context, _ *events.Message) error {
	err := h.schemaInvitations.Delete(ctx, repositoryschemainvitations.FilterExpired)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}
