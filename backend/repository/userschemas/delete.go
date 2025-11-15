package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) Delete(ctx context.Context, scopes ...repository.Scope) error {
	_, err := gorm.
		G[model.UserSchema](c.db).
		Scopes(scopes...).
		Delete(ctx)

	return err
}
