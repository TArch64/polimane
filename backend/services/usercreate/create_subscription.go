package usercreate

import (
	"context"
	"time"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (s *Service) createSubscription(ctx context.Context, tx *gorm.DB, user *model.User) error {
	return s.userSubscriptions.InsertTx(ctx, tx, &model.UserSubscription{
		UserID:         user.ID,
		PlanID:         model.SubscriptionBasic,
		TrialStartedAt: time.Now(),
		TrialEndsAt:    time.Now().Add(model.SubscriptionTrialDuration),
	})
}
