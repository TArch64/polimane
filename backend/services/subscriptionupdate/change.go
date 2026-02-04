package subscriptionupdate

import (
	"errors"

	"polimane/backend/model"
)

var (
	InvalidPlanErr = errors.New("invalid subscription plan")
)

type ChangeOptions struct {
	User *model.User
	Plan *model.SubscriptionPlan
}

func (s *Service) canUpdate(from, to *model.SubscriptionPlan) error {
	if from.IsBeta() {
		return InvalidPlanErr
	}
	if to.IsBeta() {
		return InvalidPlanErr
	}
	if from.ID == to.ID {
		return InvalidPlanErr
	}
	return nil
}
