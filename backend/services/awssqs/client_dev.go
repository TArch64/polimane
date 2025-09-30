//go:build dev

package awssqs

import (
	"github.com/aws/aws-sdk-go-v2/service/sqs"

	"polimane/backend/env"
)

func configure(environment *env.Environment) Option {
	return func(options *sqs.Options) {
		options.BaseEndpoint = &environment.AWS.SQSBaseURL
	}
}
