package migrations

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Migration func(ctx *migrationCtx) error

var migrations = []Migration{
	v0,
	v1,
}

func getVersionKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"pk": &types.AttributeValueMemberS{
			Value: "#VERSION",
		},
		"sk": &types.AttributeValueMemberS{
			Value: "#METADATA",
		},
	}
}

func getTableVersion(ctx *migrationCtx) (int, error) {
	item, err := ctx.Api.GetItem(ctx, &dynamodb.GetItemInput{
		TableName:            &ctx.TableName,
		Key:                  getVersionKey(),
		ProjectionExpression: aws.String("version"),
	})

	if err != nil {
		var notFoundErr *types.ResourceNotFoundException
		if errors.As(err, &notFoundErr) {
			return 0, nil
		}
		return 0, err
	}

	attr := item.Item["version"].(*types.AttributeValueMemberN)
	version, _ := strconv.Atoi(attr.Value)
	return version, nil
}

func Migrate(ctx_ context.Context, api *dynamodb.Client) error {
	ctx := &migrationCtx{
		Context:   ctx_,
		Api:       api,
		TableName: "polimane",
	}

	version, err := getTableVersion(ctx)
	if err != nil {
		return err
	}

	if version+1 >= len(migrations) {
		log.Printf("[DynamoDB] Table is already at version %d\n", version)
		return nil
	}

	log.Printf("[DynamoDB] Current version: %d\n", version)

	versionKey := getVersionKey()

	for i := version + 1; i < len(migrations); i++ {
		log.Printf("[DynamoDB] Running migration %d\n", i)
		err = migrations[i](ctx)
		if err != nil {
			return err
		}

		log.Printf("[DynamoDB] Migration %d complete\n", i)

		versionKey["version"] = &types.AttributeValueMemberN{
			Value: strconv.Itoa(i),
		}

		_, err = ctx.Api.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: &ctx.TableName,
			Item:      versionKey,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
