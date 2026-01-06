package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) Count(ctx context.Context, scopes ...repository.Scope) (int64, error) {
	return c.CountByColumn(ctx, "schema_id", scopes...)
}

func (c *Client) CountQuery(ctx context.Context, scopes ...repository.Scope) gorm.ChainInterface[model.UserSchema] {
	return c.CountByColumnQuery(ctx, "schema_id", scopes...)
}
