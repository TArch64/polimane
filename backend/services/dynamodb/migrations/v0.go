package migrations

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"polimane/backend/env"
)

func v0(ctx *Ctx) error {
	tags := []types.Tag{
		{
			Key:   aws.String("app"),
			Value: aws.String("polimane"),
		},
	}

	if len(env.Env().AWSAppArn) > 0 {
		tags = append(tags, types.Tag{
			Key:   aws.String("awsApplication"),
			Value: &env.Env().AWSAppArn,
		})
	}

	_, err := ctx.Api.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName:                 &ctx.TableName,
		BillingMode:               types.BillingModePayPerRequest,
		DeletionProtectionEnabled: aws.Bool(true),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("PK"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("SK"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("PK"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("SK"),
				KeyType:       types.KeyTypeRange,
			},
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String("UserNameIndex"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("SK"),
						KeyType:       types.KeyTypeHash,
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
			},
		},
		Tags: tags,
	})

	if err != nil {
		return err
	}

	return ctx.Table.Wait(ctx)
}
