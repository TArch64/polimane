package handlercleanupinvitations

import (
	"context"

	"polimane/backend/worker/events"
)

func (h *Handler) Handle(ctx context.Context, _ *events.Message) error {
	return h.schemaInvitations.DeleteExpired(ctx)
}
