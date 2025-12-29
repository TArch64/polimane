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
	Name        string
	OwnersCount int64
}

func (s *Service) logAffected(ctx context.Context, tx *gorm.DB, schemaIDs []model.ID) error {
	affected, err := s.getAffectedResources(ctx, tx, schemaIDs)
	if err != nil {
		return err
	}

	for _, resource := range affected {
		s.persistentLogger.InfoContext(ctx,
			"Schema deleted",
			slog.String("id", resource.ID.String()),
			slog.String("name", resource.Name),
			slog.Time("created_at", resource.CreatedAt),
			slog.Int64("owners_count", resource.OwnersCount),
		)
	}

	return nil
}

func (s *Service) getAffectedResources(ctx context.Context, tx *gorm.DB, schemaIDs []model.ID) ([]*affectedResource, error) {
	var resources []*affectedResource

	err := s.schemas.ListOut(ctx, &resources,
		repository.IncludeSoftDeleted,
		repository.Select(
			"MIN(id) as id",
			"MIN(name) as name",
			"MIN(schemas.created_at) as created_at",
			"COUNT(user_schemas.user_id) AS owners_count",
		),
		repository.Join("JOIN user_schemas ON schemas.id = user_schemas.schema_id"),
		repository.IDsIn(schemaIDs),
		repository.Group("id"),
	)

	if err != nil {
		return nil, err
	}

	return resources, nil
}
