package repositorybase

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"polimane/backend/repository"
)

func (c *Client[M]) Exists(ctx context.Context, scopes ...repository.Scope) (bool, error) {
	var exists bool

	scopes = append(scopes,
		repository.Select("1 AS exists"),
		repository.Limit(1),
	)

	err := c.GetOut(ctx, &exists, scopes...)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exists, nil
}
