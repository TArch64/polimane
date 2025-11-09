package schemas

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/awss3"
)

type deleteBody struct {
	IDs []model.ID `json:"ids" validate:"required"`
}

func (c *Controller) apiDelete(ctx *fiber.Ctx) error {
	var err error
	var body deleteBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	requestCtx := ctx.Context()
	err = c.userSchemas.FilterByAccess(requestCtx, user, &body.IDs, model.AccessAdmin)
	if err != nil {
		return err
	}

	err = c.schemas.DeleteMany(ctx.Context(), &repositoryschemas.DeleteOptions{
		SchemaIDs: body.IDs,
	})

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}

func (c *Controller) deleteScreenshots(ctx context.Context, schemaIDs []model.ID) error {
	objectIDs := make([]types.ObjectIdentifier, len(schemaIDs))
	for index, schemaID := range schemaIDs {
		objectIDs[index].Key = aws.String(model.SchemaScreenshotKey(schemaID))
	}

	_, err := c.s3.DeleteObjects(ctx, &s3.DeleteObjectsInput{
		Bucket: &awss3.Bucket,
		Delete: &types.Delete{Objects: objectIDs},
	})

	var notFound *types.NotFound
	if errors.As(err, &notFound) {
		return nil
	}

	return err
}
