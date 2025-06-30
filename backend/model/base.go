package model

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/guregu/dynamo/v2"
)

const IfPKExists = "attribute_exists(PK)"
const IfSKExists = "attribute_exists(SK)"
const IfKeyExists = IfPKExists + " and " + IfSKExists

type Base struct {
	PK ID  `dynamo:"PK,hash"`
	SK Key `dynamo:"SK,range"`
}

func (b *Base) PrimaryKey() PrimaryKey {
	return NewPrimaryKey(b.PK, ID(b.SK))
}

func ConditionErrToNotFound(err error) error {
	var checkFailedErr *types.ConditionalCheckFailedException
	if errors.As(err, &checkFailedErr) {
		return dynamo.ErrNotFound
	}

	return err
}
