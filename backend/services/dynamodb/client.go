package awsdynamodb

import (
	"context"

	"github.com/guregu/dynamo/v2"

	"polimane/backend/services/dynamodb/migrations"
)

var table *dynamo.Table

func Table() *dynamo.Table {
	return table
}

func Init(ctx context.Context) error {
	cfg, err := newConfig(ctx)
	if err != nil {
		return err
	}

	db := dynamo.New(*cfg, configureClient)

	err = migrations.Migrate(ctx, db)
	if err != nil {
		return err
	}

	table_ := db.Table("polimane")
	table = &table_
	return nil
}
