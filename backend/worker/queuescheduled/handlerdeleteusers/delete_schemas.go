package handlerdeleteusers

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (h *Handler) deleteUserSchemas(ctx context.Context, tx *gorm.DB, user *model.User) error {
	schemaIDs, err := h.getUserSchemaIDs(ctx, tx, user)
	if err != nil {
		return err
	}

	if len(schemaIDs) == 0 {
		return nil
	}

	if err = h.deleteSchemaRelation(ctx, tx, user, schemaIDs); err != nil {
		return err
	}

	if err = h.filterOrphanSchemaIDs(ctx, tx, &schemaIDs); err != nil {
		return err
	}

	if len(schemaIDs) == 0 {
		return nil
	}

	return h.schemaDelete.DeleteTx(ctx, tx, schemaIDs)
}

func (h *Handler) getUserSchemaIDs(ctx context.Context, tx *gorm.DB, user *model.User) ([]model.ID, error) {
	var ids []model.ID
	err := h.userSchemas.ListOutTx(ctx, tx, &ids,
		repository.IncludeSoftDeleted,
		repository.Select("schema_id"),
		repository.UserIDEq(user.ID),
	)
	return ids, err
}

func (h *Handler) deleteSchemaRelation(ctx context.Context, tx *gorm.DB, user *model.User, schemaIDs []model.ID) error {
	return h.userSchemas.DeleteTx(ctx, tx,
		repository.IncludeSoftDeleted,
		repository.UserIDEq(user.ID),
		repository.SchemaIDsIn(schemaIDs),
	)
}

type schemaIDUsedPlaceholder struct{}

func (h *Handler) filterOrphanSchemaIDs(ctx context.Context, tx *gorm.DB, schemaIDs *[]model.ID) error {
	var withUserIDs []model.ID
	err := h.userSchemas.ListOutTx(ctx, tx, &withUserIDs,
		repository.IncludeSoftDeleted,
		repository.Select("DISTINCT ON (schema_id) schema_id"),
		repository.SchemaIDsIn(*schemaIDs),
	)
	if err != nil {
		return err
	}

	var orphanIDs []model.ID
	schemaIDSet := make(map[model.ID]*schemaIDUsedPlaceholder, len(withUserIDs))
	placeholder := &schemaIDUsedPlaceholder{}
	for _, id := range withUserIDs {
		schemaIDSet[id] = placeholder
	}

	for _, id := range *schemaIDs {
		if _, exists := schemaIDSet[id]; !exists {
			orphanIDs = append(orphanIDs, id)
		}
	}

	*schemaIDs = orphanIDs
	return nil
}
