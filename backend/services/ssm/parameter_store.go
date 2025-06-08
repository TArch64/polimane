package awsssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

func GetParameter(ctx context.Context, name string) (string, error) {
	result, err := Client().GetParameter(ctx, &ssm.GetParameterInput{
		Name: &name,
	})

	if err != nil {
		return "", err
	}

	return *result.Parameter.Value, nil
}

func PutParameter(ctx context.Context, name, value string) error {
	_, err := Client().PutParameter(ctx, &ssm.PutParameterInput{
		Name:  &name,
		Value: &value,
		Type:  types.ParameterTypeString,
	})

	return err
}
