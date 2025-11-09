package users

import (
	"context"
	"strings"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type GetByEmailOptions struct {
	Email  string
	Select []string
}

func (c *Client) GeyByEmail(ctx context.Context, options *GetByEmailOptions) (*model.User, error) {
	query := gorm.
		G[*model.User](c.db).
		Where("email = ?", options.Email)

	if len(options.Select) > 0 {
		query = query.Select(strings.Join(options.Select, ", "))
	}

	return query.Take(ctx)
}
