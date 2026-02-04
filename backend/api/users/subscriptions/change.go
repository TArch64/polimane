package subscriptions

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/services/subscriptionupdate"
)

type ChangeBody struct {
	PlanID model.SubscriptionPlanID `json:"planId" validate:"oneof=basic pro"`
}

func (c *Controller) Change(ctx *fiber.Ctx) (err error) {
	var body ChangeBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	toPlan := model.Plans[body.PlanID]
	reqCtx := ctx.Context()

	changeOptions := &subscriptionupdate.ChangeOptions{
		User: user,
		Plan: toPlan,
	}

	if user.Subscription.Plan().Tier > toPlan.Tier {
		err = c.subscriptionUpdate.Downgrade(reqCtx, changeOptions)
	} else {
		err = c.subscriptionUpdate.Upgrade(reqCtx, changeOptions)
	}

	if errors.Is(err, subscriptionupdate.InvalidPlanErr) {
		return fiber.ErrBadRequest
	}
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
