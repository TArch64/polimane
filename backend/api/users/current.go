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
	Plan     model.SubscriptionPlan      `json:"plan"`
	Status   model.SubscriptionStatus    `json:"status"`
	Counters *model.SubscriptionCounters `json:"counters"`
	Limits   *model.SubscriptionLimits   `json:"limits"`
}

func (c *Controller) Current(ctx *fiber.Ctx) error {
	session := auth.GetSession(ctx)

	return ctx.JSON(UserResponse{
		ID:            session.User.ID,
		FirstName:     session.User.FirstName,
		LastName:      session.User.LastName,
		Email:         session.User.Email,
		EmailVerified: session.WorkosUser.EmailVerified,

		Subscription: &SubscriptionResponse{
			Plan:     session.User.Subscription.Plan,
			Status:   session.User.Subscription.Status,
			Counters: session.User.Subscription.Counters.Data(),
			Limits:   session.User.Subscription.Limits(),
		},
	})
}
