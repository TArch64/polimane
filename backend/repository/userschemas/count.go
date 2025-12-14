package userschemas

import (
	"context"

	"polimane/backend/repository"
)

func (c *Client) Count(ctx context.Context, scopes ...repository.Scope) (int64, error) {
	return c.CountByColumn(ctx, "schema_id", scopes...)
}
