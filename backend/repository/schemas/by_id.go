package schemas

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
)

type ByIDOptions struct {
	Ctx      context.Context
	User     *model.User
	SchemaID modelbase.ID
	Select   []string
}

func (c *Client) ByID(options *ByIDOptions) (*model.Schema, error) {
	var err error

	err = c.userSchemas.HasAccess(options.Ctx, options.User.ID, options.SchemaID)
	if err != nil {
		return nil, err
	}

	var schema model.Schema
	query := c.db.WithContext(options.Ctx)

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	err = query.Take(&schema, options.SchemaID).Error
	return &schema, err
}
