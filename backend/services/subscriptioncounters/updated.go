package subscriptioncounters

import (
	"polimane/backend/model"
)

type updatedCounter struct {
	ID    model.ID
	Count uint16
}
