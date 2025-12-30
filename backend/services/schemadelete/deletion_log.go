package schemadelete

import (
	"context"
	"log/slog"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

type affectedResource struct {
	*model.Identifiable
	*model.Timestamps
	Name           string
	InitiatorID    *model.ID
	InitiatorEmail *string
	UsersCount     int64
}

func (s *Service) logAffected(ctx context.Context, tx *gorm.DB, schemaIDs []model.ID) error {
	affected, err := s.getAffectedResources(ctx, tx, schemaIDs)
	if err != nil {
		return err
	}

	for _, resource := range affected {
		s.logAffectedResource(ctx, resource)
	}

	return nil
}

func (s *Service) getAffectedResources(ctx context.Context, tx *gorm.DB, schemaIDs []model.ID) ([]*affectedResource, error) {
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
		repository.Join("JOIN user_schemas ON schemas.id = user_schemas.schema_id"),
		repository.Join("LEFT JOIN users ON schemas.deleted_by = users.id"),
		repository.IDsIn(schemaIDs, "schemas"),
		repository.Group("schemas.id"),
	)

	if err != nil {
		return nil, err
	}

	return resources, nil
}

func (s *Service) logAffectedResource(ctx context.Context, resource *affectedResource) {
	args := []any{
		slog.String("id", resource.ID.String()),
		slog.String("name", resource.Name),
		slog.Time("created_at", resource.CreatedAt),
	}

	if resource.InitiatorID == nil || !resource.InitiatorID.Valid {
		args = append(args,
			slog.String("initiator_id", "system"),
			slog.String("initiator_email", "system"),
		)
	} else {
		args = append(args,
			slog.String("initiator_id", resource.InitiatorID.String()),
			slog.String("initiator_email", *resource.InitiatorEmail),
		)
	}

	args = append(args, slog.Int64("users_count", resource.UsersCount))
	s.persistentLogger.InfoContext(ctx, "Schema deleted", args...)
}
