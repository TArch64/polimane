package users

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"github.com/workos/workos-go/v4/pkg/workos_errors"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/services/workos"
)

var (
	ErrEmailVerificationCodeExpired = base.NewReasonedError(fiber.StatusBadRequest, "CodeExpired")
)

type BodyEmailVerify struct {
	Code string `json:"code" validate:"required,numeric,len=6"`
}

func (c *Controller) EmailVerify(ctx *fiber.Ctx) (err error) {
	var body BodyEmailVerify
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)

	_, err = c.workos.UserManagement.VerifyEmail(ctx.Context(), usermanagement.VerifyEmailOpts{
		User: user.WorkosID,
		Code: body.Code,
	})

	var httpError *workos_errors.HTTPError
	if errors.As(err, &httpError) && workos.GetErrorCode(httpError) == workos.CodeEmailVerificationCodeExpired {
		return ErrEmailVerificationCodeExpired
	}

	if err != nil {
		return err
	}

	c.signals.InvalidateWorkosUserCache.Emit(ctx.Context(), user.WorkosID)
	return base.NewSuccessResponse(ctx)
}
