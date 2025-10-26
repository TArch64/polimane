package schemas

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"gorm.io/gorm"

	"polimane/backend/model"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/awss3"
)

type DeleteOptions struct {
	User     *model.User
	SchemaID model.ID
}

func (c *Client) Delete(ctx context.Context, options *DeleteOptions) (err error) {
	err = c.userSchemas.HasAccess(ctx, options.User.ID, options.SchemaID, model.AccessAdmin)
	if err != nil {
		return err
	}

	err = c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err = tx.Delete(&model.Schema{}, options.SchemaID).Error
		if err != nil {
			return err
		}

		err = c.userSchemas.DeleteTx(ctx, tx, &repositoryuserschemas.DeleteOptions{
			UserID:   options.User.ID,
			SchemaID: options.SchemaID,
		})

		if err != nil {
			return err
		}

		return c.deleteScreenshot(ctx, options.SchemaID)
	})

	if err != nil {
		return err
	}

	c.signals.InvalidateUserCache.Emit(ctx, options.User.ID)
	return nil
}

func (c *Client) deleteScreenshot(ctx context.Context, schemaID model.ID) error {
	key := model.SchemaScreenshotKey(schemaID)

	_, err := c.s3.DeleteObject(ctx, &s3.DeleteObjectInput{
		Key:    &key,
		Bucket: &awss3.Bucket,
	})

	var notFound *types.NotFound
	if errors.As(err, &notFound) {
		return nil
	}

	return err
}
