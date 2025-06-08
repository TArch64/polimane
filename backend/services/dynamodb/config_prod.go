//go:build !dev

package awsdynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const TableName = "polimane-prod"
const TableLockParameter = "/polimane/prod/db/lock"

func configureClient(options *dynamodb.Options) {}
