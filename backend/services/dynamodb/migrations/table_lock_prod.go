//go:build !dev

package migrations

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/ssm/types"

	dynamodbconfig "polimane/backend/services/dynamodb/config"
	awsssm "polimane/backend/services/ssm"
)

func isTableLocked(ctx context.Context) (bool, error) {
	locked, err := awsssm.GetParameter(ctx, dynamodbconfig.TableLockParameter)
	var notFoundErr *types.ParameterNotFound
	if errors.As(err, &notFoundErr) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return locked == "true", nil
}

func setTableLock(ctx context.Context, isLocked bool) {
	value := strconv.FormatBool(isLocked)
	fmt.Printf("[DynamoDB] Setting table lock to %s\n", value)
	_ = awsssm.PutParameter(ctx, dynamodbconfig.TableLockParameter, value)
}
