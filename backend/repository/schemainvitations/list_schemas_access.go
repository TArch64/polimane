package schemainvitations

import (
	"context"
	"fmt"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) ListSchemasAccessOut(ctx context.Context, schemaIDs []model.ID, out interface{}) error {
	return c.ListOut(ctx, out,
		repository.Select(
			"email",
			"MIN(access) AS access",
			"MIN(access) = MAX(access) AS is_even_access",
			fmt.Sprintf("COUNT(schema_id) = %d AS is_all_access", len(schemaIDs)),
		),
		repository.SchemaIDsIn(schemaIDs),
		repository.Group("email"),
		repository.Order("email"),
	)
}
