package schemas

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/services/awss3"
)

type DeleteOptions struct {
	SchemaIDs []model.ID
}

func (c *Client) DeleteMany(ctx context.Context, options *DeleteOptions) error {
	_, err := gorm.
		G[model.Schema](c.db).
		Where("id IN (?)", options.SchemaIDs).
		Delete(ctx)

	if err != nil {
		return err
	}

	return c.deleteScreenshots(ctx, options.SchemaIDs)
}

func (c *Client) deleteScreenshots(ctx context.Context, schemaIDs []model.ID) error {
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
