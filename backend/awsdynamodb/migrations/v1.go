package migrations

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func v1(ctx *migrationCtx) error {
	_, err := ctx.Api.UpdateTable(ctx, &dynamodb.UpdateTableInput{
		TableName: &ctx.TableName,
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("sk"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		GlobalSecondaryIndexUpdates: []types.GlobalSecondaryIndexUpdate{
			{
				Create: &types.CreateGlobalSecondaryIndexAction{
					IndexName: aws.String("EmailIndex"),
					KeySchema: []types.KeySchemaElement{
						{
							AttributeName: aws.String("sk"),
							KeyType:       types.KeyTypeHash,
						},
					},
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeAll,
					},
				},
			},
		},
	})

	return err
}
