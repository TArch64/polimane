package awss3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Option = func(*s3.Options)

type Client interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...Option) (*s3.PutObjectOutput, error)
}

func Provider(config *aws.Config) Client {
	return s3.NewFromConfig(*config)
}
