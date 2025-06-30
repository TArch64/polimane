package migrations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/guregu/dynamo/v2"
)

type Ctx struct {
	context.Context
	DB        *dynamo.DB
	Table     *dynamo.Table
	Api       *dynamodb.Client
	TableArn  string
	TableName string
}
