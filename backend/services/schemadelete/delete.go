package schemadelete

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (s *Service) Delete(ctx context.Context, IDs []model.ID) (err error) {
	return s.DeleteTx(ctx, s.schemas.DB, IDs)
}

func (s *Service) DeleteTx(ctx context.Context, tx *gorm.DB, IDs []model.ID) (err error) {
	if err = s.screenshot.Delete(ctx, IDs); err != nil {
		return err
	}

	return s.schemas.DeleteTx(ctx, tx,
		repository.IncludeSoftDeleted,
		repository.IDsIn(IDs),
	)
}
