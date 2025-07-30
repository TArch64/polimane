package signal

import (
	"github.com/maniartech/signals"

	"polimane/backend/model/modelbase"
)

type Container struct {
	InvalidateUserCache       signals.Signal[modelbase.ID]
	InvalidateWorkosUserCache signals.Signal[string]
	InvalidateAuthCache       signals.Signal[string]
}

func Provider() *Container {
	return &Container{
		InvalidateUserCache:       signals.New[modelbase.ID](),
		InvalidateWorkosUserCache: signals.New[string](),
		InvalidateAuthCache:       signals.New[string](),
	}
}
