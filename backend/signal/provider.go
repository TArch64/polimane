package signal

import (
	"github.com/maniartech/signals"

	"polimane/backend/model/modelbase"
)

type Container struct {
	InvalidateUserCache signals.Signal[modelbase.ID]
	InvalidateAuthCache signals.Signal[string]
}

func Provider() *Container {
	return &Container{
		InvalidateUserCache: signals.New[modelbase.ID](),
		InvalidateAuthCache: signals.New[string](),
	}
}
