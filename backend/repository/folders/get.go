package folders

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) Get(ctx context.Context, scopes ...repository.Scope) (*model.Folder, error) {
	var out model.Folder
	err := c.GetOut(ctx, &out, scopes...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetOut(ctx context.Context, out interface{}, scopes ...repository.Scope) error {
	return gorm.
		G[model.Folder](c.db).
		Scopes(scopes...).
		Scan(ctx, out)
}
