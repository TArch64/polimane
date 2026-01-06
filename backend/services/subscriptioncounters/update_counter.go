package subscriptioncounters

import (
	"context"

	"gorm.io/datatypes"

	"polimane/backend/model"
	"polimane/backend/repository"
)

type setCountersFunc func(expr *datatypes.JSONSetExpression) *datatypes.JSONSetExpression

func (s *Service) updateCounter(ctx context.Context, userID model.ID, setCounters setCountersFunc) error {
	err := s.userSubscriptions.UpdateColumn(ctx,
		"counters",
		setCounters(datatypes.JSONSet("counters")),
		repository.UserIDEq(userID),
	)

	if err != nil {
		return err
	}

	s.signals.InvalidateUserCache.Emit(ctx, userID)
	return nil
}
