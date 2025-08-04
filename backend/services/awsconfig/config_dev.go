//go:build dev

package awsconfig

import (
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"

	"polimane/backend/env"
)

var S3Bucket = "polimane-dev"

func configure(e *env.Environment, options *config.LoadOptions) error {
	options.DefaultRegion = e.AWS.Region
	options.Credentials = credentials.NewStaticCredentialsProvider(e.AWS.AccessKeyID, e.AWS.SecretAccessKey, "")
	return nil
}
