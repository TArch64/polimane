package signal

import (
	"polimane/backend/model"
)

type UpdateUserCacheEvent struct {
	UserID model.ID
	Update func(user *model.User)
}

func NewUpdateUserCacheEvent(userID model.ID, update func(user *model.User)) *UpdateUserCacheEvent {
	return &UpdateUserCacheEvent{
		UserID: userID,
		Update: update,
	}
}
