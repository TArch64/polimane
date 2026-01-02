package handlerdeleteusers

import (
	"context"

	"polimane/backend/model"
)

type Deleter interface {
	Collect(ctx context.Context, user *model.User) error
	Delete(ctx context.Context) error
	LogResults(ctx context.Context)
	Cleanup()
}
