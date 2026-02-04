package subscriptions

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/model"
)

func (c *Controller) Plans(ctx *fiber.Ctx) error {
	user := auth.GetSessionUser(ctx)

	if user.Subscription.PlanID == model.SubscriptionBeta {
		return ctx.JSON([]*model.SubscriptionPlan{})
	}

	return ctx.JSON([]*model.SubscriptionPlan{
		model.BasicPlan,
		model.ProPlan,
	})
}
