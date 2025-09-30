package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/views"
	"polimane/backend/views/templates"
)

func (c *Controller) apiPreview(ctx *fiber.Ctx) error {
	schemaId, err := base.GetParamID(ctx, schemaIdParam)
	if err != nil {
		return err
	}

	schema, err := c.schemas.ByID(&repositoryschemas.ByIDOptions{
		Ctx:      ctx.Context(),
		User:     auth.GetSessionUser(ctx),
		SchemaID: schemaId,
	})

	if err != nil {
		return err
	}

	content, err := c.renderer.Render(&views.RenderOptions{
		View: views.TemplateSchemaPreview,
		Data: templates.NewSchemaPreviewData(schema),
	})

	if err != nil {
		return err
	}

	return ctx.Type("svg").SendString(content)
}
