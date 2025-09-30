package queuedebounced

import (
	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
)

type Queue struct {
	*queue.Base
}

func Provider() queue.Interface {
	q := &Queue{
		Base: &queue.Base{},
	}

	q.HandleEvent(events.EventSchemaScreenshot, q.ProcessSchemaScreenshot)
	return q
}

func (q *Queue) Name() string {
	return events.QueueDebounced
}

var _ queue.Interface = (*Queue)(nil)
