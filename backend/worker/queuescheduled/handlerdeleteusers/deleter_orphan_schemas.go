package handlerdeleteusers

import (
	"context"
	"log/slog"
	"strings"

	"polimane/backend/model"
	"polimane/backend/repository"
	repositoryschemas "polimane/backend/repository/schemas"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/logpersistent"
	"polimane/backend/services/schemadelete"
)

type schemaIDUsedPlaceholder struct{}

type DeleterOrphanSchemas struct {
	user             *model.User
	affected         []model.ID
	Schemas          *repositoryschemas.Client
	UserSchemas      *repositoryuserschemas.Client
	SchemaDelete     *schemadelete.Service
	PersistentLogger *logpersistent.Logger
}

func (d *DeleterOrphanSchemas) Collect(ctx context.Context, user *model.User) error {
	d.user = user

	return d.UserSchemas.ListOut(ctx, &d.affected,
		repository.IncludeSoftDeleted,
		repository.Select("schema_id"),
		repository.UserIDEq(user.ID),
	)
}

func (d *DeleterOrphanSchemas) Delete(ctx context.Context) error {
	if len(d.affected) == 0 {
		return nil
	}
	return d.SchemaDelete.Delete(ctx, d.affected)
}

func (d *DeleterOrphanSchemas) filterOrphanSchemaIDs(ctx context.Context) error {
	var withUserIDs []model.ID
	err := d.UserSchemas.ListOut(ctx, &withUserIDs,
		repository.IncludeSoftDeleted,
		repository.Select("DISTINCT ON (schema_id) schema_id"),
		repository.SchemaIDsIn(d.affected),
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

	for _, id := range d.affected {
		if _, exists := schemaIDSet[id]; !exists {
			orphanIDs = append(orphanIDs, id)
		}
	}

	d.affected = orphanIDs
	return nil
}

func (d *DeleterOrphanSchemas) LogResults(ctx context.Context) {
	if len(d.affected) == 0 {
		d.PersistentLogger.InfoContext(ctx, "no orphan Schemas to delete for user",
			slog.String("operation", "user_deletion"),
		)
	} else {
		var schemaIDsStrBuilder strings.Builder
		schemaIDsStrBuilder.Grow(len(d.affected))
		for index, id := range d.affected {
			if index > 0 {
				schemaIDsStrBuilder.WriteRune(',')
			}
			schemaIDsStrBuilder.WriteString(id.String())
		}

		d.PersistentLogger.InfoContext(ctx, "deleted orphan Schemas due to user deletion",
			slog.String("operation", "user_deletion"),
			slog.Int("schemas_count", len(d.affected)),
			slog.String("schema_ids", schemaIDsStrBuilder.String()),
		)
	}
}

func (d *DeleterOrphanSchemas) Cleanup() {
	d.user = nil
	d.affected = nil
}
