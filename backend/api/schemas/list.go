package schemas

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
)

type listQuery struct {
	Offset uint16 `query:"offset" validate:"gte=0,lte=65535"`
	Limit  uint8  `query:"limit" validate:"gte=1,lte=100"`
}

type listResponse struct {
	Folders []*listFolder `json:"folders"`
	Schemas []*listSchema `json:"schemas"`
	Total   int64         `json:"total"`
}

type listContext struct {
	context.Context
	user         *model.User
	query        *listQuery
	schemasTotal int64
	foldersTotal int64
	res          *listResponse
}

func (c *Controller) apiList(ctx *fiber.Ctx) (err error) {
	var query listQuery
	if err = base.ParseQuery(ctx, &query); err != nil {
		return err
	}

	eg, egCtx := errgroup.WithContext(ctx.Context())

	listCtx := &listContext{
		Context: egCtx,
		query:   &query,
		res:     &listResponse{},
		user:    auth.GetSessionUser(ctx),
	}

	eg.Go(func() (err error) {
		if err = c.queryFolders(listCtx); err != nil {
			return err
		}

		return c.querySchemas(listCtx)
	})

	eg.Go(func() error {
		return c.countFolders(listCtx)
	})

	eg.Go(func() error {
		return c.countSchemas(listCtx)
	})

	if err = eg.Wait(); err != nil {
		return err
	}

	listCtx.res.Total = listCtx.schemasTotal + listCtx.foldersTotal
	return ctx.JSON(listCtx.res)
}
