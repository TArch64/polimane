package migrations

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/guregu/dynamo/v2"

	"polimane/backend/model"
)

func getTableVersion(ctx *Ctx) (*model.Version, error) {
	var version model.Version

	err := ctx.Table.
		Get("PK", model.PKVersion).
		Range("SK", dynamo.Equal, model.SKVersion).
		One(ctx, &version)

	var notFoundErr *types.ResourceNotFoundException
	if errors.As(err, &notFoundErr) {
		return model.NewVersion(), nil
	}
	if errors.Is(err, dynamo.ErrNotFound) {
		return model.IntVersion(0), nil
	}
	if err != nil {
		return nil, err
	}

	return &version, nil
}

func setTableVersion(ctx *Ctx, version *model.Version, index int) error {
	version.Version = index
	return ctx.Table.Put(version).Run(ctx)
}
