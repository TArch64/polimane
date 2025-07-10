package signal

import (
	"github.com/maniartech/signals"

	"polimane/backend/model/modelbase"
)

var InvalidateAuthCache = signals.New[modelbase.ID]()
