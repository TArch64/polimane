package awsdynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/guregu/dynamo/v2"

	"polimane/backend/services/dynamodb/migrations"
	awsssm "polimane/backend/services/ssm"
)

var table *dynamo.Table

func Table() *dynamo.Table {
	return table
}

func newConfig(ctx context.Context) (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	return &cfg, err
}

func isTableLocked(ctx context.Context) (bool, error) {
	locked, err := awsssm.GetParameter(ctx, TableLockParameter)
	if err != nil {
		return false, err
	}

	return locked == "true", nil
}

func setTableLock(ctx context.Context, isLocked bool) {
	var value string
	if isLocked {
		value = "true"
	} else {
		value = "false"
	}

	_ = awsssm.PutParameter(ctx, TableLockParameter, value)
}

func Init(ctx context.Context) error {
	cfg, err := newConfig(ctx)
	if err != nil {
		return err
	}

	db := dynamo.New(*cfg, configureClient)

	isLocked, err := isTableLocked(ctx)
	if err != nil {
		return err
	}

	if !isLocked {
		setTableLock(ctx, true)
		err = migrations.Migrate(ctx, db)
		setTableLock(ctx, false)
		if err != nil {
			return err
		}
	}

	table_ := db.Table(TableName)
	table = &table_
	return nil
}
