package subscriptioncounters

import (
	"context"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (s *Service) SyncSchemasCreated(ctx context.Context, userID model.ID) error {
	return s.updateCounter(ctx, userID, func(expr *datatypes.JSONSetExpression) *datatypes.JSONSetExpression {
		countQuery := s.userSchemas.CountQuery(ctx,
			repository.UserIDEq(userID),
		)

		return expr.Set("{schemasCreated}", gorm.Expr("to_jsonb((?))", countQuery))
	})
}
