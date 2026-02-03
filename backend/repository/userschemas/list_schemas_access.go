package userschemas

import (
	"context"
	"fmt"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) ListSchemasAccessOut(ctx context.Context, schemaIDs []model.ID, out interface{}) error {
	return c.ListOut(ctx, out,
		repository.Select(
			"users.id",
			"email",
			"first_name",
			"last_name",
			"MIN(access) AS access",
			"MIN(access) = MAX(access) AS is_even_access",
			fmt.Sprintf("COUNT(schema_id) = %d AS is_all_access", len(schemaIDs)),
			"MIN(user_schemas.created_at) AS created_at",
		),
		IncludeUsersScope,
		repository.SchemaIDsIn(schemaIDs),
		repository.Group("users.id, email, first_name, last_name"),
		repository.Order("created_at"),
		repository.Limit(repository.DefaultBatch),
	)
}
