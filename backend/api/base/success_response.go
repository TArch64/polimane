package base

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Success bool `json:"success"`
}

func NewSuccessResponse() *SuccessResponse {
	return &SuccessResponse{Success: true}
}

func (r *SuccessResponse) AsJSON(ctx *fiber.Ctx) error {
	return ctx.JSON(r)
}
