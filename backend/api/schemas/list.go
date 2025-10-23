package schemas

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

type listItem struct {
	ID              model.ID          `json:"id"`
	Name            string            `json:"name"`
	Access          model.AccessLevel `json:"access"`
	BackgroundColor string            `json:"backgroundColor"`
	ScreenshotedAt  *time.Time        `json:"screenshotedAt"`
	ScreenshotPath  *string           `json:"screenshotPath"`
}

func newListItem(schema *model.Schema) *listItem {
	return &listItem{
		ID:              schema.ID,
		Name:            schema.Name,
		Access:          schema.Access,
		BackgroundColor: schema.BackgroundColor,
		ScreenshotedAt:  schema.ScreenshotedAt,
		ScreenshotPath:  schema.ScreenshotPath(),
	}
}

func (c *Controller) apiList(ctx *fiber.Ctx) error {
	schemas, err := c.schemas.ByUser(ctx.Context(), &repositoryschemas.ByUserOptions{
		User: auth.GetSessionUser(ctx),
		Select: []string{
			"id",
			"name",
			"screenshoted_at",
			"background_color",
			"user_schemas.access AS access",
		},
	})

	if err != nil {
		return err
	}

	if schemas == nil {
		return ctx.JSON([]*listItem{})
	}

	items := make([]*listItem, len(schemas))
	for i, schema := range schemas {
		items[i] = newListItem(schema)
	}

	return ctx.JSON(items)
}
