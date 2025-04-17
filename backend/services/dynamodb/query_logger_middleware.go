package awsdynamodb

import (
	"context"
	"fmt"
	"log"
	"maps"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
)

type queryLoggerMiddleware struct{}

var _ middleware.InitializeMiddleware = (*queryLoggerMiddleware)(nil)

func (m *queryLoggerMiddleware) ID() string {
	return "QueryLoggerMiddleware"
}

func (m *queryLoggerMiddleware) HandleInitialize(
	ctx context.Context,
	in middleware.InitializeInput,
	next middleware.InitializeHandler,
) (
	out middleware.InitializeOutput,
	metadata middleware.Metadata,
	err error,
) {
	log.Println(m.formatMessage(in))
	return next.HandleInitialize(ctx, in)
}

func (m *queryLoggerMiddleware) formatMessage(in middleware.InitializeInput) string {
	var building string

	switch typed := in.Parameters.(type) {
	case *dynamodb.QueryInput:
		building = fmt.Sprintf("[DynamoDB/Query]: ")
		if typed.KeyConditionExpression != nil {
			building += fmt.Sprintf("Where %s", *typed.KeyConditionExpression)
		} else if len(typed.KeyConditions) > 0 {
			building += fmt.Sprintf("Where %s", m.formatConditionMap(typed.KeyConditions))
		}
		if typed.IndexName != nil {
			building += " By Index " + *typed.IndexName
		}
		if typed.ProjectionExpression != nil {
			building += " Select " + m.formatProjectionExpression(*typed.ProjectionExpression, typed.ExpressionAttributeNames)
		}
		return building

	case *dynamodb.GetItemInput:
		building = fmt.Sprintf("[DynamoDB/GetItem]: Where %s", m.formatAttrMap(typed.Key))
		if typed.ProjectionExpression != nil {
			building += " Select " + m.formatProjectionExpression(*typed.ProjectionExpression, typed.ExpressionAttributeNames)
		}
		return building

	case *dynamodb.PutItemInput:
		building = fmt.Sprintf("[DynamoDB/PutItem]: Attrs %s", m.formatAttrMap(typed.Item))
		if typed.ConditionExpression != nil {
			building += " If " + m.formatConditionExpression(*typed.ConditionExpression)
		}
		return building

	case *dynamodb.DeleteItemInput:
		building = fmt.Sprintf("[DynamoDB/DeleteItem]: Where %s", m.formatAttrMap(typed.Key))
		if typed.ConditionExpression != nil {
			building += " If " + m.formatConditionExpression(*typed.ConditionExpression)
		}
		return building

	case *dynamodb.UpdateItemInput:
		building = fmt.Sprintf("[DynamoDB/UpdateItem]: Where %s %s",
			m.formatAttrMap(typed.Key),
			m.formatUpdateExpression(*typed.UpdateExpression, typed.ExpressionAttributeValues),
		)
		if typed.ConditionExpression != nil {
			building += " If " + m.formatConditionExpression(*typed.ConditionExpression)
		}
		return building
	default:
		return ""
	}
}

func (m *queryLoggerMiddleware) formatAttrMap(attrMap map[string]types.AttributeValue) string {
	attrMap = reorderQueryAttrMap(attrMap)
	var result string

	for name, attr := range attrMap {
		if result != "" {
			result += ", "
		}

		result += m.formatCondition(name, types.ComparisonOperatorEq, attr)
	}

	return result
}

func (m *queryLoggerMiddleware) formatConditionMap(conditionMap map[string]types.Condition) string {
	conditionMap = reorderQueryAttrMap(conditionMap)
	var result string

	for name, condition := range conditionMap {
		if result != "" {
			result += ", "
		}

		result += m.formatCondition(name, condition.ComparisonOperator, condition.AttributeValueList...)
	}

	return result
}

func (m *queryLoggerMiddleware) formatCondition(name string, operator types.ComparisonOperator, values ...types.AttributeValue) string {
	result := fmt.Sprintf("%s %s", name, m.formatConditionOperator(operator))

	if len(values) > 0 {
		result += " " + m.formatSliceValue(values)
	}

	return result
}

func (m *queryLoggerMiddleware) formatConditionOperator(operator types.ComparisonOperator) string {
	switch operator {
	case types.ComparisonOperatorEq:
		return "="
	case types.ComparisonOperatorNe:
		return "!="
	case types.ComparisonOperatorLe:
		return "<="
	case types.ComparisonOperatorLt:
		return "<"
	case types.ComparisonOperatorGe:
		return ">="
	case types.ComparisonOperatorGt:
		return ">"
	case types.ComparisonOperatorBeginsWith:
		return "begins_with"
	case types.ComparisonOperatorBetween:
		return "between"
	case types.ComparisonOperatorIn:
		return "in"
	case types.ComparisonOperatorNotNull:
		return "not_null"
	case types.ComparisonOperatorNull:
		return "null"
	case types.ComparisonOperatorContains:
		return "contains"
	case types.ComparisonOperatorNotContains:
		return "not_contains"
	default:
		return string(operator)
	}
}

func (m *queryLoggerMiddleware) formatSliceValue(values []types.AttributeValue) string {
	if len(values) == 1 {
		return m.formatAttrValue(values[0])
	}

	result := ""
	for i, value := range values {
		if i > 0 {
			result += ", "
		}
		result += m.formatAttrValue(value)
	}
	return "(" + result + ")"
}

func (m *queryLoggerMiddleware) formatProjectionExpression(expression string, attrNames map[string]string) string {
	for name, value := range attrNames {
		expression = strings.Replace(expression, name, value, 1)
	}
	return expression
}

func (m *queryLoggerMiddleware) formatConditionExpression(expression string) string {
	return strings.TrimSuffix(strings.TrimPrefix(expression, "("), ")")
}

func (m *queryLoggerMiddleware) formatUpdateExpression(expression string, attrValues map[string]types.AttributeValue) string {
	for name, value := range attrValues {
		expression = strings.ReplaceAll(expression, name, m.formatAttrValue(value))
	}
	return strings.Replace(expression, "SET", "Set", 1)
}

func (m *queryLoggerMiddleware) formatAttrValue(value types.AttributeValue) string {
	switch typed := value.(type) {
	case *types.AttributeValueMemberS:
		return "\"" + typed.Value + "\""
	case *types.AttributeValueMemberN:
		return typed.Value
	case *types.AttributeValueMemberM:
		return "{TRUNCATED MAP}"
	default:
		log.Println("[QueryLoggerMiddleware]: Unknown attribute type:", reflect.TypeOf(value))
		return ""
	}
}

var priorityQueryAttrOrder = []string{"PK", "SK"}

func reorderQueryAttrMap[V any](input map[string]V) (dest map[string]V) {
	temp := maps.Clone(input)
	dest = make(map[string]V)

	for _, name := range priorityQueryAttrOrder {
		if value, ok := temp[name]; ok {
			dest[name] = value
			delete(temp, name)
		}
	}

	maps.Copy(dest, temp)
	return dest
}
