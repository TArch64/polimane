package schemas

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

type ListDeletedQuery struct {
	Offset uint16 `query:"offset" validate:"gte=0,lte=65535"`
	Limit  uint8  `query:"limit" validate:"gte=1,lte=100"`
}

type ListDeletedResponse struct {
	Schemas []*ListSchema `json:"schemas"`
	Total   int64         `json:"total"`
}

func (c *Controller) ListDeleted(ctx *fiber.Ctx) (err error) {
	var query ListDeletedQuery
	if err = base.ParseQuery(ctx, &query); err != nil {
		return err
	}

	currentUser := auth.GetSessionUser(ctx)
	listResponse := &ListDeletedResponse{}

	if query.Offset == 0 {
		eg, egCtx := errgroup.WithContext(ctx.Context())

		eg.Go(func() error {
			return c.queryDeletedSchemas(egCtx, currentUser, &query, listResponse)
		})

		eg.Go(func() error {
			return c.countDeletedSchemas(egCtx, listResponse, currentUser)
		})

		if err = eg.Wait(); err != nil {
			return err
		}
	} else {
		err = c.queryDeletedSchemas(ctx.Context(), currentUser, &query, listResponse)
		if err != nil {
			return err
		}
	}

	return ctx.JSON(listResponse)
}

func (c *Controller) queryDeletedSchemas(
	ctx context.Context,
	user *model.User,
	query *ListDeletedQuery,
	res *ListDeletedResponse,
) (err error) {
	scopes := c.deletedSchemasFilter(user)

	scopes = append(scopes,
		repository.Select(
			"id",
			"name",
			"screenshoted_at",
			"background_color",
			"user_schemas.access",
		),
		repository.Paginate(query.Offset, query.Limit),
		repository.Order("user_schemas.deleted_at DESC"),
	)

	if err = c.userSchemas.ListOut(ctx, &res.Schemas, scopes...); err != nil {
		return err
	}

	if res.Schemas == nil {
		res.Schemas = []*ListSchema{}
	}

	return nil
}

func (c *Controller) countDeletedSchemas(ctx context.Context, res *ListDeletedResponse, user *model.User) (err error) {
	res.Total, err = c.userSchemas.Count(ctx, c.deletedSchemasFilter(user)...)
	return err
}

func (c *Controller) deletedSchemasFilter(user *model.User) []repository.Scope {
	return []repository.Scope{
		repository.SoftDeletedOnly("user_schemas"),
		repository.UserIDEq(user.ID),
		repositoryuserschemas.IncludeSchemasScope(),
	}
}
