package schemas

import (
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/awsconfig"
)

type apiUpdateScreenshotBody struct {
	Src string `json:"src" validate:"required"`
}

func (c *Controller) apiUpdateScreenshot(ctx *fiber.Ctx) error {
	schemaId, err := base.GetParamID(ctx, schemaIdParam)
	if err != nil {
		return err
	}

	var body apiUpdateScreenshotBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	if err = c.uploadScreenshot(ctx.Context(), schemaId, body.Src); err != nil {
		return err
	}

	screenshotedAt := time.Now()

	err = c.schemas.Update(&repositoryschemas.UpdateOptions{
		Ctx:      ctx.Context(),
		User:     auth.GetSessionUser(ctx),
		SchemaID: schemaId,
		Updates:  &model.Schema{ScreenshotedAt: &screenshotedAt},
	})

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}

func (c *Controller) uploadScreenshot(ctx context.Context, schemaId model.ID, src string) error {
	key := model.SchemaScreenshotKey(schemaId)

	_, err := c.s3.PutObject(ctx, &s3.PutObjectInput{
		Key:         &key,
		Bucket:      &awsconfig.S3Bucket,
		ACL:         types.ObjectCannedACLPrivate,
		Body:        strings.NewReader(src),
		ContentType: aws.String("image/svg+xml"),
	})

	return err
}
