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

const ParamUserID = "userID"
const ParamDefUserID = ":" + ParamUserID

type ControllerOptions struct {
	fx.In
	Users             *repositoryusers.Client
	UserSchemas       *repositoryuserschemas.Client
	SchemaInvitations *repositoryschemainvitations.Client
	Workos            *workos.Client
}

type Controller struct {
	users             *repositoryusers.Client
	userSchemas       *repositoryuserschemas.Client
	schemaInvitations *repositoryschemainvitations.Client
	workos            *workos.Client
}

func Provider(options ControllerOptions) *Controller {
	return &Controller{
		users:             options.Users,
		userSchemas:       options.UserSchemas,
		schemaInvitations: options.SchemaInvitations,
		workos:            options.Workos,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "users", func(group fiber.Router) {
		group.Get("", c.List)
		group.Post("", c.Add)

		base.WithGroup(group, "invitations", func(group fiber.Router) {
			group.Delete("", c.DeleteInvitation)
			group.Patch("access", c.UpdateInvitationAccess)
		})

		base.WithGroup(group, ParamDefUserID, func(group fiber.Router) {
			group.Delete("", c.Delete)
			group.Patch("access", c.UpdateAccess)
		})
	})
}

func (c *Controller) FilterSchemaIDsByAccess(ctx *fiber.Ctx, IDs *[]model.ID) error {
	return c.userSchemas.FilterByAccess(ctx.Context(), auth.GetSessionUser(ctx), IDs, model.AccessAdmin)
}

type BulkOperationBody struct {
	IDs []model.ID `json:"ids" validate:"required"`
}

type InvitationBody struct {
	BulkOperationBody
	Email string `json:"email" validate:"required,email,max=255"`
}
