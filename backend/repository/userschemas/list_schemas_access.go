package userschemas

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) ListSchemasAccessOut(ctx context.Context, schemaIDs []model.ID, out interface{}) error {
	const userColumns = "users.id, email, first_name, last_name"
	const accessColumns = "MIN(access) AS access, MIN(access) != MAX(access) AS is_uneven_access"
	const createdAtColumn = "MIN(user_schemas.created_at) AS created_at"
	const selectExpr = userColumns + "," + accessColumns + "," + createdAtColumn

	return c.DB.
		WithContext(ctx).
		Select(selectExpr).
		Table("user_schemas").
		Scopes(IncludeUsersLegacyScope).
		Where("schema_id IN (?)", schemaIDs).
		Group(userColumns).
		Order("created_at").
		Limit(repository.DefaultBatch).
		Scan(out).
		Error
}
