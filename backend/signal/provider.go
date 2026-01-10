package signal

import (
	"github.com/maniartech/signals"

	"polimane/backend/model"
)

type Container struct {
	InvalidateUserCache       signals.Signal[model.ID]
	UpdateUserCacheSync       *signals.SyncSignal[*UpdateUserCacheEvent]
	InvalidateWorkosUserCache signals.Signal[string]
	InvalidateAuthCache       signals.Signal[string]
}

func Provider() *Container {
	return &Container{
		InvalidateUserCache:       signals.New[model.ID](),
		UpdateUserCacheSync:       signals.NewSync[*UpdateUserCacheEvent](),
		InvalidateWorkosUserCache: signals.New[string](),
		InvalidateAuthCache:       signals.New[string](),
	}
}
