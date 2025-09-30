package queue

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func ParseBody[B any](message *types.Message, dest *B) error {
	return json.Unmarshal([]byte(*message.Body), dest)
}
