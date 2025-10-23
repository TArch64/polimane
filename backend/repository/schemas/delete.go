package schemas

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/services/awsconfig"
)

type DeleteOptions struct {
	User     *model.User
	SchemaID model.ID
}

func (i *Client) Delete(ctx context.Context, options *DeleteOptions) (err error) {
	err = i.userSchemas.HasAccess(ctx, options.User.ID, options.SchemaID, model.AccessAdmin)
	if err != nil {
		return err
	}

	err = i.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(&model.Schema{}, options.SchemaID).Error; err != nil {
			return err
		}

		if err = i.userSchemas.DeleteTx(tx, options.User.ID, options.SchemaID); err != nil {
			return err
		}

		return i.deleteScreenshot(ctx, options.SchemaID)
	})

	if err != nil {
		return err
	}

	i.signals.InvalidateUserCache.Emit(ctx, options.User.ID)
	return nil
}

func (i *Client) deleteScreenshot(ctx context.Context, schemaId model.ID) error {
	key := model.SchemaScreenshotKey(schemaId)

	_, err := i.s3.DeleteObject(ctx, &s3.DeleteObjectInput{
		Key:    &key,
		Bucket: &awsconfig.S3Bucket,
	})

	var notFound *types.NotFound
	if errors.As(err, &notFound) {
		return nil
	}

	return err
}
