package schemas

import (
	"bytes"
	"context"
	"encoding/base64"
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
	Src string `json:"src" validate:"required,url,startswith=data:image/webp;base64"`
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
	dataBase64 := src[strings.IndexByte(src, ',')+1:]
	data, err := base64.StdEncoding.DecodeString(dataBase64)
	if err != nil {
		return err
	}

	_, err = c.s3.PutObject(ctx, &s3.PutObjectInput{
		Key:         &key,
		Bucket:      &awsconfig.S3Bucket,
		ACL:         types.ObjectCannedACLPrivate,
		Body:        bytes.NewReader(data),
		ContentType: aws.String("image/webp"),
	})

	return err
}
