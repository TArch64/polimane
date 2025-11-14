package schemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) ListOut(ctx context.Context, out interface{}, scopes ...repository.Scope) error {
	err := gorm.
		G[model.Schema](c.db).
		Scopes(scopes...).
		Scan(ctx, out)

	if err != nil {
		return err
	}

	return repository.DoAfterScan(out)
}
