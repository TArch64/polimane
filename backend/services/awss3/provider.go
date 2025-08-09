package awss3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/fx"
)

type Option = func(*s3.Options)

type ClientOptions struct {
	fx.In
	Config *aws.Config
}

type Client interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...Option) (*s3.PutObjectOutput, error)
	DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...Option) (*s3.DeleteObjectOutput, error)
}

func Provider(options ClientOptions) Client {
	return s3.NewFromConfig(*options.Config)
}
