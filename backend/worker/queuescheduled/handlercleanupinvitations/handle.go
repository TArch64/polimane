package handlercleanupinvitations

import (
	"context"
	"errors"
	"log/slog"

	"gorm.io/gorm"

	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	"polimane/backend/worker/events"
)

func (h *Handler) Handle(ctx context.Context, _ *events.Message) error {
	affected, err := h.schemaInvitations.DeleteCounted(ctx, repositoryschemainvitations.FilterExpired)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		return err
	}

	h.logAffected(ctx, affected)
	return nil
}

func (h *Handler) logAffected(ctx context.Context, affected int) {
	h.persistentLogger.InfoContext(ctx, "cleaned up expired schema invitations",
		slog.String("operation", "cleanup_expired_invitations"),
		slog.Int("affected", affected),
	)
}
