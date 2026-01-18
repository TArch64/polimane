package repositorybase

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"polimane/backend/repository"
)

func (c *Client[M]) Exists(ctx context.Context, scopes ...repository.Scope) (exists bool, err error) {
	scopes = append(scopes,
		repository.Select("1 AS exists"),
		repository.Limit(1),
	)

	err = c.GetOut(ctx, &exists, scopes...)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		exists = false
		err = nil
	}

	return
}
