package subscriptionupdate

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (s *Service) Upgrade(ctx context.Context, options *ChangeOptions) (err error) {
	from := options.User.Subscription.Plan()

	if err = s.canUpgrade(from, options.Plan); err != nil {
		return err
	}

	err = s.userSubscriptions.Update(ctx,
		model.UserSubscription{PlanID: options.Plan.ID},
		repository.UserIDEq(options.User.ID),
	)
	if err != nil {
		return err
	}

	s.signals.InvalidateUserCache.Emit(ctx, options.User.ID)
	return nil
}

func (s *Service) canUpgrade(from, to *model.SubscriptionPlan) error {
	if err := s.canUpdate(from, to); err != nil {
		return err
	}

	if from.Tier > to.Tier {
		return InvalidPlanErr
	}

	return nil
}
