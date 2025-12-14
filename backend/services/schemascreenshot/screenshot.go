package schemascreenshot

import (
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"polimane/backend/model"
	"polimane/backend/repository"
	"polimane/backend/services/awss3"
	"polimane/backend/views"
	"polimane/backend/views/templates"
)

type ScreenshotOptions struct {
	Schema *model.Schema
}

func (s *Service) Screenshot(ctx context.Context, options *ScreenshotOptions) error {
	content, err := s.renderer.Render(&views.RenderOptions{
		View:   views.TemplateSchemaPreview,
		Data:   templates.NewSchemaPreviewData(options.Schema),
		Minify: true,
	})

	if err != nil {
		return err
	}

	key := model.SchemaScreenshotKey(options.Schema.ID)

	_, err = s.s3.PutObject(ctx, &s3.PutObjectInput{
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

	return s.schemas.Update(ctx,
		model.Schema{ScreenshotedAt: &screenshotedAt},
		repository.IDEq(options.Schema.ID),
	)
}
