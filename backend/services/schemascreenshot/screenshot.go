package schemascreenshot

import (
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/awsconfig"
	"polimane/backend/views"
	"polimane/backend/views/templates"
)

type ScreenshotOptions struct {
	Schema *model.Schema
}

func (i *Impl) Screenshot(ctx context.Context, options *ScreenshotOptions) error {
	content, err := i.renderer.Render(&views.RenderOptions{
		View:   views.TemplateSchemaPreview,
		Data:   templates.NewSchemaPreviewData(options.Schema),
		Minify: true,
	})

	if err != nil {
		return err
	}

	key := options.Schema.ScreenshotKey()

	_, err = i.s3.PutObject(ctx, &s3.PutObjectInput{
		Key:         &key,
		Bucket:      &awsconfig.S3Bucket,
		ACL:         types.ObjectCannedACLPrivate,
		Body:        strings.NewReader(content),
		ContentType: aws.String("image/svg+xml"),
	})

	if err != nil {
		return err
	}

	screenshotedAt := time.Now()

	return i.schemas.Update(&repositoryschemas.UpdateOptions{
		Ctx:      ctx,
		SchemaID: options.Schema.ID,
		Updates:  &model.Schema{ScreenshotedAt: &screenshotedAt},
	})
}
