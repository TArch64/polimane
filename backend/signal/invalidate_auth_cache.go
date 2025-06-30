package signal

import (
	"github.com/maniartech/signals"

	"polimane/backend/model"
)

var InvalidateAuthCache = signals.New[model.ID]()
