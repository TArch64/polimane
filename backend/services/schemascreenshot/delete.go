package schemascreenshot

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"polimane/backend/model"
	"polimane/backend/services/awss3"
)

func (s *Service) Delete(ctx context.Context, schemaIDs []model.ID) error {
	objectIDs := make([]types.ObjectIdentifier, len(schemaIDs))
	for index, schemaID := range schemaIDs {
		objectIDs[index].Key = aws.String(model.SchemaScreenshotKey(schemaID))
	}

	_, err := s.s3.DeleteObjects(ctx, &s3.DeleteObjectsInput{
		Bucket: &awss3.Bucket,
		Delete: &types.Delete{Objects: objectIDs},
	})

	var notFound *types.NotFound
	if errors.As(err, &notFound) {
		return nil
	}

	return err
}
