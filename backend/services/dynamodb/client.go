package awsdynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/guregu/dynamo/v2"

	"polimane/backend/base"
	dynamodbconfig "polimane/backend/services/dynamodb/config"
	"polimane/backend/services/dynamodb/migrations"
)

var table *dynamo.Table

func Table() *dynamo.Table {
	return table
}

func newConfig(ctx context.Context) (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	return &cfg, err
}

func Init(ctx context.Context) error {
	cfg, err := newConfig(ctx)
	if err != nil {
		return base.TagError("dynamodb.config", err)
	}

	db := dynamo.New(*cfg, dynamodbconfig.ConfigureClient)
	table_ := db.Table(dynamodbconfig.TableName)
	table = &table_

	migrationCtx := &migrations.Ctx{
		Context:   ctx,
		Api:       db.Client().(*dynamodb.Client),
		Table:     table,
		TableName: table.Name(),
	}

	err = migrations.Migrate(migrationCtx)
	return base.TagError("dynamodb.migrations", err)
}
