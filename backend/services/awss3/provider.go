package awss3

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/fx"
)

type Option = func(*s3.Options)

type ClientOptions struct {
	fx.In
	Config *aws.Config
}

func Provider(options ClientOptions) *s3.Client {
	return s3.NewFromConfig(*options.Config)
}
