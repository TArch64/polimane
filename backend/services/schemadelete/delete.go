package schemadelete

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (s *Service) Delete(ctx context.Context, IDs []model.ID) (err error) {
	affected, err := s.getAffectedResources(ctx, IDs)
	if err != nil {
		return err
	}

	if err = s.screenshot.Delete(ctx, IDs); err != nil {
		return err
	}

	err = s.schemas.Delete(ctx,
		repository.HardDelete,
		repository.IDsIn(IDs),
	)
	if err != nil {
		return err
	}

	s.logAffectedResources(ctx, affected)
	return nil
}
