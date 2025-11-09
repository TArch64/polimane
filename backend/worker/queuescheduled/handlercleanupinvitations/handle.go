package handlercleanupinvitations

import (
	"context"

	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	"polimane/backend/worker/events"
)

func (h *Handler) Handle(ctx context.Context, _ *events.Message) error {
	return h.schemaInvitations.DeleteMany(ctx, repositoryschemainvitations.FilterExpired)
}
