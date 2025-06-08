//go:build dev

package awsdynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/smithy-go/middleware"
)

const TableName = "polimane-dev"
const TableLockParameter = "/polimane/dev/db/lock"

func configureClient(options *dynamodb.Options) {
	options.BaseEndpoint = aws.String("http://dynamodb:8000")

	options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
		return stack.Initialize.Add(&queryLoggerMiddleware{}, middleware.After)
	})
}
