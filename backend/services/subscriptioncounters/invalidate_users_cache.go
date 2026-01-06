package subscriptioncounters

import (
	"context"

	"polimane/backend/model"
)

func (s *Service) invalidateUsersCache(ctx context.Context, userIDs []model.ID) {
	for _, userID := range userIDs {
		s.signals.InvalidateUserCache.Emit(ctx, userID)
	}
}
