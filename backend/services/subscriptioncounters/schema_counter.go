package subscriptioncounters

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"polimane/backend/model"
)

const changeSchemaCounterSQL = `
UPDATE user_schemas
SET counters = jsonb_increment(counters, @counter_name, @delta)
WHERE schema_id = @schema_id
RETURNING (counters->>@counter_name)::smallint AS count
`

type schemaCounterDeps struct {
	db *gorm.DB
}

type schemaCounterOptions struct {
	*schemaCounterDeps
	name     string
	localSet model.Set[*model.UserSchema, uint16]
}

type SchemaCounter struct {
	*schemaCounterOptions
}

func newSchemaCounter(options *schemaCounterOptions) *SchemaCounter {
	return &SchemaCounter{schemaCounterOptions: options}
}

func (s *SchemaCounter) AddTx(ctx context.Context, tx *gorm.DB, userSchema *model.UserSchema, value int) error {
	return s.change(ctx, tx, userSchema, value)
}

func (s *SchemaCounter) RemoveTx(ctx context.Context, tx *gorm.DB, userSchema *model.UserSchema, value int) error {
	return s.change(ctx, tx, userSchema, -value)
}

func (s *SchemaCounter) change(ctx context.Context, tx *gorm.DB, userSchema *model.UserSchema, value int) error {
	var updated []*updatedCounter

	err := gorm.
		G[model.UserSubscription](tx).
		Raw(changeSchemaCounterSQL,
			sql.Named("counter_name", s.name),
			sql.Named("delta", value),
			sql.Named("schema_id", userSchema.SchemaID),
		).
		Scan(ctx, &updated)

	if err != nil {
		return err
	}

	if len(updated) > 0 {
		s.localSet(userSchema, updated[0].Count)
	}

	return nil
}
