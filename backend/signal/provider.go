package signal

import (
	"github.com/maniartech/signals"

	"polimane/backend/model"
)

type Container struct {
	InvalidateUserCache       signals.Signal[model.ID]
	InvalidateWorkosUserCache signals.Signal[string]
	InvalidateAuthCache       signals.Signal[string]
}

func Provider() *Container {
	return &Container{
		InvalidateUserCache:       signals.New[model.ID](),
		InvalidateWorkosUserCache: signals.New[string](),
		InvalidateAuthCache:       signals.New[string](),
	}
}
