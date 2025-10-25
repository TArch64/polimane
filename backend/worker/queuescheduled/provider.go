package queuescheduled

import (
	"go.uber.org/fx"

	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
	"polimane/backend/worker/queuescheduled/handlercleanupinvitations"
)

type Queue struct {
	*queue.Base
	handlerCleanupInvitations *handlercleanupinvitations.Handler
}

type ProviderOptions struct {
	fx.In
	HandlerCleanupInvitations *handlercleanupinvitations.Handler
}

func Provider(options ProviderOptions) queue.Interface {
	q := &Queue{
		Base:                      queue.NewBase(),
		handlerCleanupInvitations: options.HandlerCleanupInvitations,
	}

	q.HandleEvent(events.EventCleanupInvitations, options.HandlerCleanupInvitations.Handle)
	return q
}

func (q *Queue) Name() string {
	return events.QueueScheduled
}
