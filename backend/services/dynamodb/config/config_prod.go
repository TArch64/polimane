//go:build !dev

package dynamodbconfig

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

const TableName = "polimane-prod"
const TableLockParameter = "/polimane/prod/db/lock"

func ConfigureClient(options *dynamodb.Options) {}
