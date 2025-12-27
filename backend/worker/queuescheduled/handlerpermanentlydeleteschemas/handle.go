package handlerpermanentlydeleteschemas

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/repository"
	"polimane/backend/worker/events"
)

func (h *Handler) Handle(ctx context.Context, _ *events.Message) error {
	schemaIDs, err := h.getSoftDeletedSchemaIDs(ctx)
	if err != nil {
		return err
	}

	return h.delete.Delete(ctx, schemaIDs)
}

func (h *Handler) getSoftDeletedSchemaIDs(ctx context.Context) ([]model.ID, error) {
	var IDs []model.ID

	err := h.schemas.ListOut(ctx, &IDs,
		repository.Select("id"),
		repository.SoftDeletedDaysAgo(30),
	)

	if err != nil {
		return nil, err
	}

	return IDs, nil
}
