package migrations

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"polimane/backend/model"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/guregu/dynamo/v2"
)

type Migration func(ctx *migrationCtx) error

var migrations = []Migration{
	v0,
	v1,
}

func getTableVersion(ctx *migrationCtx) (*model.Version, error) {
	var version model.Version

	err := ctx.Table.
		Get("PK", model.PKVersion).
		Range("SK", dynamo.Equal, model.SKVersion).
		One(ctx, &version)

	var notFoundErr *types.ResourceNotFoundException
	if errors.As(err, &notFoundErr) {
		return model.NewVersion(), nil
	}
	if err != nil {
		return nil, err
	}

	return &version, nil
}

func Migrate(ctx_ context.Context, db *dynamo.DB) error {
	table := db.Table("polimane")
	ctx := &migrationCtx{
		Context:   ctx_,
		Api:       db.Client().(*dynamodb.Client),
		Table:     &table,
		TableName: table.Name(),
	}

	version, err := getTableVersion(ctx)
	if err != nil {
		return err
	}

	if version.IsLatest(len(migrations)) {
		log.Printf("[DynamoDB] Table is already at version %d\n", version.Version)
		return nil
	}

	log.Printf("[DynamoDB] Current version: %d\n", version.Version)

	startVersion := version.NextVersion()

	for i := startVersion; i < len(migrations); i++ {
		log.Printf("[DynamoDB] Running migration %d\n", i)
		err = migrations[i](ctx)
		if err != nil {
			return err
		}

		log.Printf("[DynamoDB] Migration %d complete\n", i)

		version.Version = i
		err = ctx.Table.Put(version).Run(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}
