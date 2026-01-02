package schemadelete

import (
	"context"
	"log/slog"

	"polimane/backend/model"
	"polimane/backend/repository"
	"polimane/backend/services/logpersistent"
)

type affectedResource struct {
	*model.Identifiable
	*model.Timestamps
	Name           string
	InitiatorID    *model.ID
	InitiatorEmail *string
	UsersCount     int64
}

func (a *affectedResource) getAttrs() []any {
	args := []any{
		logpersistent.OperationSchemaDeletion,
		slog.String("id", a.ID.String()),
		slog.String("name", a.Name),
		slog.Time("created_at", a.CreatedAt),
	}

	if a.InitiatorID == nil || !a.InitiatorID.Valid {
		args = append(args,
			slog.String("initiator_id", "system"),
			slog.String("initiator_email", "system"),
		)
	} else {
		args = append(args,
			slog.String("initiator_id", a.InitiatorID.String()),
			slog.String("initiator_email", *a.InitiatorEmail),
		)
	}

	return append(args, slog.Int64("users_count", a.UsersCount))
}

func (s *Service) getAffectedResources(ctx context.Context, schemaIDs []model.ID) ([]*affectedResource, error) {
	var resources []*affectedResource

	err := s.schemas.ListOut(ctx, &resources,
		repository.IncludeSoftDeleted,
		repository.Select(
			"MIN(schemas.id) as id",
			"MIN(name) as name",
			"MIN(schemas.created_at) as created_at",
			"MIN(users.id) AS initiator_id",
			"MIN(users.email) AS initiator_email",
			"COUNT(user_schemas.user_id) AS users_count",
		),
		repository.Join("LEFT JOIN user_schemas ON schemas.id = user_schemas.schema_id"),
		repository.Join("LEFT JOIN users ON schemas.deleted_by = users.id"),
		repository.IDsIn(schemaIDs, "schemas"),
		repository.Group("schemas.id"),
	)

	if err != nil {
		return nil, err
	}

	return resources, nil
}

func (s *Service) logAffectedResources(ctx context.Context, affected []*affectedResource) {
	for _, resource := range affected {
		s.persistentLogger.InfoContext(ctx, "schema deleted", resource.getAttrs()...)
	}
}
