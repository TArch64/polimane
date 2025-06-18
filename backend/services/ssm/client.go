package awsssm

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var client *ssm.Client
var mu sync.RWMutex

func newConfig(ctx context.Context) (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	return &cfg, err
}

func Init(ctx context.Context) error {
	mu.Lock()
	defer mu.Unlock()

	cfg, err := newConfig(ctx)
	if err != nil {
		return err
	}

	client = ssm.NewFromConfig(*cfg)
	return nil
}

func Client() *ssm.Client {
	mu.RLock()
	defer mu.RUnlock()
	return client
}
