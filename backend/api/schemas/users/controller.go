package users

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	repositoryusers "polimane/backend/repository/users"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/workos"
)

const userIDParam = "userID"

type ControllerOptions struct {
	fx.In
	Users             *repositoryusers.Client
	UserSchemas       *repositoryuserschemas.Client
	SchemaInvitations *repositoryschemainvitations.Client
	WorkosClient      *workos.Client
}

type Controller struct {
	users             *repositoryusers.Client
	userSchemas       *repositoryuserschemas.Client
	schemaInvitations *repositoryschemainvitations.Client
	workosClient      *workos.Client
}

func Provider(options ControllerOptions) *Controller {
	return &Controller{
		users:             options.Users,
		userSchemas:       options.UserSchemas,
		schemaInvitations: options.SchemaInvitations,
		workosClient:      options.WorkosClient,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "users", func(group fiber.Router) {
		group.Get("", c.apiList)
		group.Post("", c.apiAdd)

		base.WithGroup(group, "invitations", func(group fiber.Router) {
			group.Delete("", c.apiDeleteInvitation)
			group.Patch("access", c.apiUpdateInvitationAccess)
		})

		base.WithGroup(group, ":"+userIDParam, func(group fiber.Router) {
			group.Delete("", c.apiDelete)
			group.Patch("access", c.apiUpdateAccess)
		})
	})
}

func (c *Controller) FilterSchemaIDsByAccess(ctx *fiber.Ctx, IDs *[]model.ID) error {
	err := c.userSchemas.FilterByAccess(ctx.Context(), auth.GetSessionUser(ctx), IDs, model.AccessAdmin)
	if err != nil {
		return err
	}
	if len(*IDs) == 0 {
		return fiber.ErrBadRequest
	}
	return nil
}

type bulkOperationBody struct {
	IDs []model.ID `json:"ids" validate:"required"`
}

type invitationBody struct {
	bulkOperationBody
	Email string `json:"email" validate:"required,email,max=255"`
}
