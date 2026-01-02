package awscloudwatch

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"go.uber.org/fx"
)

const (
	base = "/polimane-prod"
)

var (
	GroupPersistent = base + "/persistent"
)

type ClientOptions struct {
	fx.In
	Config *aws.Config
}

func Provider(options ClientOptions) *cloudwatchlogs.Client {
	return cloudwatchlogs.NewFromConfig(*options.Config)
}
