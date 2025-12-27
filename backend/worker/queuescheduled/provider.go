package queuescheduled

import (
	"go.uber.org/fx"

	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
	"polimane/backend/worker/queuescheduled/handlercleanupinvitations"
	"polimane/backend/worker/queuescheduled/handlerdeleteusers"
	"polimane/backend/worker/queuescheduled/handlerpermanentlydeleteschemas"
)

type Queue struct {
	*queue.Base
}

type ProviderOptions struct {
	fx.In
	HandlerCleanupInvitations       *handlercleanupinvitations.Handler
	HandlerDeleteUsers              *handlerdeleteusers.Handler
	HandlerPermanentlyDeleteSchemas *handlerpermanentlydeleteschemas.Handler
}

func Provider(options ProviderOptions) queue.Interface {
	q := &Queue{Base: queue.NewBase()}
	q.HandleEvent(events.EventCleanupInvitations, options.HandlerCleanupInvitations.Handle)
	q.HandleEvent(events.EventDeleteUsers, options.HandlerDeleteUsers.Handle)
	q.HandleEvent(events.EventPermanentlyDeleteSchemas, options.HandlerPermanentlyDeleteSchemas.Handle)
	return q
}

func (q *Queue) Name() string {
	return events.QueueScheduled
}
