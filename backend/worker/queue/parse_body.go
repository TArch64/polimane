package queue

import (
	"encoding/json"

	"polimane/backend/worker/events"
)

func ParseBody[B any](message *events.Message, dest *B) error {
	return json.Unmarshal([]byte(message.Body), dest)
}
