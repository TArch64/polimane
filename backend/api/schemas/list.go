package schemas

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

type listQuery struct {
	Offset uint16 `query:"offset" validate:"gte=0,lte=65535"`
	Limit  uint8  `query:"limit" validate:"gte=1,lte=100"`
}

type listResponse struct {
	List  []*listItem `json:"list"`
	Total int64       `json:"total"`
}

type listItem struct {
	ID              model.ID          `json:"id"`
	Name            string            `json:"name"`
	Access          model.AccessLevel `json:"access"`
	BackgroundColor string            `json:"backgroundColor"`
	ScreenshotedAt  *time.Time        `json:"screenshotedAt"`
	ScreenshotPath  *string           `json:"screenshotPath"`
}

func newListItem(schema *model.Schema, access model.AccessLevel) *listItem {
	return &listItem{
		ID:              schema.ID,
		Name:            schema.Name,
		Access:          access,
		BackgroundColor: schema.BackgroundColor,
		ScreenshotedAt:  schema.ScreenshotedAt,
		ScreenshotPath:  model.SchemaScreenshotPath(schema.ID, schema.ScreenshotedAt),
	}
}

func (l *listItem) AfterFind(_ *gorm.DB) error {
	l.ScreenshotPath = model.SchemaScreenshotPath(l.ID, l.ScreenshotedAt)
	return nil
}

func (c *Controller) apiList(ctx *fiber.Ctx) error {
	var query listQuery
	if err := base.ParseQuery(ctx, &query); err != nil {
		return err
	}

	eg := errgroup.Group{}
	res := &listResponse{}
	user := auth.GetSessionUser(ctx)

	eg.Go(func() error {
		return c.queryList(ctx.Context(), user, &query, res)
	})

	eg.Go(func() error {
		return c.countList(ctx.Context(), user, res)
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	return ctx.JSON(res)
}

func (c *Controller) queryList(
	ctx context.Context,
	user *model.User,
	query *listQuery,
	res *listResponse,
) error {
	return c.schemas.ListByUserOut(ctx, &repositoryschemas.ListByUserOptions{
		User: user,

		Pagination: &model.Pagination{
			Offset: query.Offset,
			Limit:  query.Limit,
		},

		Select: []string{
			"id",
			"name",
			"screenshoted_at",
			"background_color",
			"user_schemas.access AS access",
		},
	}, &res.List)
}

func (c *Controller) countList(ctx context.Context, user *model.User, res *listResponse) error {
	count, err := c.schemas.CountByUser(ctx, &repositoryschemas.CountByUserOptions{
		User: user,
	})

	if err != nil {
		return err
	}

	res.Total = count
	return nil
}
