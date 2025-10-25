package handlercleanupinvitations

import (
	"context"

	"polimane/backend/worker/events"
)

func (q *Handler) Handle(ctx context.Context, _ *events.Message) error {
	println("Cleaning up invitations...")
	return nil
}
