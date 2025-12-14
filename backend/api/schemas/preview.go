package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/repository"
	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/views"
	"polimane/backend/views/templates"
)

func (c *Controller) Preview(ctx *fiber.Ctx) error {
	schemaID, err := base.GetParamID(ctx, ParamSchemaID)
	if err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	schema, err := c.schemas.Get(ctx.Context(),
		repository.IDEq(schemaID),
		repositoryschemas.IncludeUserSchemaScope(user.ID),
	)

	if err != nil {
		return err
	}

	content, err := c.renderer.Render(&views.RenderOptions{
		View:   views.TemplateSchemaPreview,
		Data:   templates.NewSchemaPreviewData(schema),
		Minify: true,
	})

	if err != nil {
		return err
	}

	return ctx.Type("svg").SendString(content)
}
