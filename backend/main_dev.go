//go:build dev

package main

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"polimane/backend/api"
	"polimane/backend/awsdynamodb"
)

func createTable(ctx context.Context) error {
	_, err := awsdynamodb.Client().CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: aws.String("polimane"),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("pk"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("sk"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("pk"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("sk"),
				KeyType:       types.KeyTypeRange,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	})

	if err != nil {
		var resourceInUseErr *types.ResourceInUseException
		if errors.As(err, &resourceInUseErr) {
			return nil
		}
		return err
	}

	return nil
}

func initDynamoDB() error {
	ctx := context.Background()
	err := awsdynamodb.Init(ctx)
	if err != nil {
		return err
	}

	return createTable(ctx)
}

func main() {
	var err error
	if err = initDynamoDB(); err != nil {
		panic(err)
	}

	err = api.New().Listen(":3000")
	if err != nil {
		panic(err)
	}
}
