package subscriptioncounters

import (
	"context"

	"gorm.io/datatypes"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (s *Service) SyncSchemasCreated(ctx context.Context, userID model.ID) error {
	count, err := s.userSchemas.Count(ctx,
		repository.UserIDEq(userID),
	)
	if err != nil {
		return err
	}

	return s.updateCounter(ctx, userID, func(expr *datatypes.JSONSetExpression) *datatypes.JSONSetExpression {
		return expr.Set("{schemasCreated}", count)
	})
}
