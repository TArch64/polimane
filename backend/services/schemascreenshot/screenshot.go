package schemascreenshot

import (
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/awss3"
	"polimane/backend/views"
	"polimane/backend/views/templates"
)

type ScreenshotOptions struct {
	Schema *model.Schema
}

func (i *Service) Screenshot(ctx context.Context, options *ScreenshotOptions) error {
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
		Bucket:      &awss3.Bucket,
		ACL:         awss3.ObjectACL,
		Body:        strings.NewReader(content),
		ContentType: aws.String("image/svg+xml"),
	})

	if err != nil {
		return err
	}

	screenshotedAt := time.Now()

	return i.schemas.Update(ctx, &repositoryschemas.UpdateOptions{
		SchemaID: options.Schema.ID,
		Updates:  &model.Schema{ScreenshotedAt: &screenshotedAt},
	})
}
