package schemas

import (
	"time"

	"polimane/backend/model"
	"polimane/backend/repository"
	repositoryschemas "polimane/backend/repository/schemas"
)

type listSchema struct {
	ID              model.ID          `json:"id"`
	Name            string            `json:"name"`
	Access          model.AccessLevel `json:"access"`
	BackgroundColor string            `json:"backgroundColor"`
	ScreenshotedAt  *time.Time        `json:"screenshotedAt"`
	ScreenshotPath  *string           `json:"screenshotPath"`
}

func newListSchema(schema *model.Schema, access model.AccessLevel) *listSchema {
	return &listSchema{
		ID:              schema.ID,
		Name:            schema.Name,
		Access:          access,
		BackgroundColor: schema.BackgroundColor,
		ScreenshotedAt:  schema.ScreenshotedAt,
		ScreenshotPath:  model.SchemaScreenshotPath(schema.ID, schema.ScreenshotedAt),
	}
}

func (l *listSchema) AfterScan() error {
	l.ScreenshotPath = model.SchemaScreenshotPath(l.ID, l.ScreenshotedAt)
	return nil
}

func (c *Controller) querySchemas(ctx *listContext) (err error) {
	limit := ctx.query.Limit - uint8(len(ctx.res.Folders))
	if limit == 0 {
		ctx.res.Schemas = []*listSchema{}
		return nil
	}

	err = c.schemas.ListOut(ctx, &ctx.res.Schemas,
		repository.Select(
			"id",
			"name",
			"screenshoted_at",
			"background_color",
			"user_schemas.access",
		),
		repositoryschemas.IncludeUserSchemaScope(ctx.user.ID),
		repository.Paginate(ctx.query.Offset, limit),
		repository.Order("schemas.created_at DESC"),
		repository.DoAfterScan,
	)

	if err != nil {
		return err
	}

	if ctx.res.Schemas == nil {
		ctx.res.Schemas = []*listSchema{}
	}

	return nil
}

func (c *Controller) countSchemas(ctx *listContext) (err error) {
	ctx.schemasTotal, err = c.schemas.Count(ctx,
		repositoryschemas.IncludeUserSchemaScope(ctx.user.ID),
	)
	return err
}
