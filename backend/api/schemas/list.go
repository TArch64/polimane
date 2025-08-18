package schemas

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

type listItem struct {
	ID             model.ID   `json:"id"`
	Name           string     `json:"name"`
	ScreenshotedAt *time.Time `json:"screenshotedAt"`
	ScreenshotPath *string    `json:"screenshotPath"`
}

func newListItem(schema *model.Schema) *listItem {
	return &listItem{
		ID:             schema.ID,
		Name:           schema.Name,
		ScreenshotedAt: schema.ScreenshotedAt,
		ScreenshotPath: schema.ScreenshotPath(),
	}
}

func (c *Controller) apiList(ctx *fiber.Ctx) error {
	schemas, err := c.schemas.ByUser(&repositoryschemas.ByUserOptions{
		Ctx:    ctx.Context(),
		User:   auth.GetSessionUser(ctx),
		Select: []string{"id", "name", "screenshoted_at"},
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
