package migrations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type migrationCtx struct {
	context.Context
	Api       *dynamodb.Client
	TableArn  string
	TableName string
}
