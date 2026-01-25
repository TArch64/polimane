package authfactors

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/api/base"
	"polimane/backend/services/workos"
)

const ParamFactorID = "factorID"
const ParamDefFactorID = ":" + ParamFactorID

type ControllerOptions struct {
	fx.In
	Workos *workos.Client
}

type Controller struct {
	workos *workos.Client
}

func Provider(options ControllerOptions) *Controller {
	return &Controller{
		workos: options.Workos,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "auth-factors", func(group fiber.Router) {
		group.Get("", c.ListAuthFactors)
		group.Post("", c.AuthFactorCreate)
		group.Post("init", c.AuthFactorsInit)
		group.Delete(ParamDefFactorID, c.AuthFactorDelete)
	})
}
