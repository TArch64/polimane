package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/model"
)

type UserResponse struct {
	ID            model.ID              `json:"id"`
	FirstName     string                `json:"firstName"`
	LastName      string                `json:"lastName"`
	Email         string                `json:"email"`
	EmailVerified bool                  `json:"isEmailVerified"`
	Subscription  *SubscriptionResponse `json:"subscription"`
}

type SubscriptionResponse struct {
	Status   model.SubscriptionStatus    `json:"status"`
	Counters *model.SubscriptionCounters `json:"counters"`
	Plan     *PlanResponse               `json:"plan"`
}

type PlanResponse struct {
	ID     model.SubscriptionPlanID  `json:"id"`
	Tier   uint8                     `json:"tier"`
	Limits *model.SubscriptionLimits `json:"limits"`
}

func (c *Controller) Current(ctx *fiber.Ctx) error {
	session := auth.GetSession(ctx)
	plan := session.User.Subscription.Plan()

	return ctx.JSON(UserResponse{
		ID:            session.User.ID,
		FirstName:     session.User.FirstName,
		LastName:      session.User.LastName,
		Email:         session.User.Email,
		EmailVerified: session.WorkosUser.EmailVerified,

		Subscription: &SubscriptionResponse{
			Status:   session.User.Subscription.Status,
			Counters: session.User.Subscription.Counters.Data(),

			Plan: &PlanResponse{
				ID:     plan.ID,
				Tier:   plan.Tier,
				Limits: plan.Limits,
			},
		},
	})
}
