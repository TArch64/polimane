package subscriptioncounters

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"polimane/backend/model"
)

const changeSchemaCounterSQL = `
UPDATE user_schemas
SET counters = jsonb_increment(counters, @counter_name, TO_JSONB(@counter_delta))
WHERE schema_id = @schema_id
RETURNING (counters->>@counter_name)::smallint AS count
`

const setSchemaCounterSQL = `
UPDATE user_schemas
SET counters = jsonb_set(counters, string_to_array(@counter_name, '.'), TO_JSONB(@counter_value))
WHERE schema_id = @schema_id
`

type schemaCounterDeps struct {
	db *gorm.DB
}

type schemaCounterOptions[CV counterValue, CD counterDelta] struct {
	*schemaCounterDeps
	name         string
	counterValue *model.Accessor[*model.UserSchema, CV]
	counterLimit *model.Accessor[*model.UserSubscription, *CV]
}

type SchemaCounter[CV counterValue, CD counterDelta] struct {
	*schemaCounterOptions[CV, CD]
}

func newSchemaCounter[CV counterValue, CD counterDelta](options *schemaCounterOptions[CV, CD]) *SchemaCounter[CV, CD] {
	return &SchemaCounter[CV, CD]{schemaCounterOptions: options}
}

func (s *SchemaCounter[CV, CD]) CanAdd(user *model.User, data *model.UserSchema, delta CV) bool {
	limit := s.counterLimit.Get(user.Subscription)
	if limit == nil {
		return true
	}
	value := s.counterValue.Get(data) + delta
	return value <= *limit
}

func (s *SchemaCounter[CV, CD]) CanSet(user *model.User, value CV) bool {
	if limit := s.counterLimit.Get(user.Subscription); limit != nil {
		return value <= *limit
	}

	return true
}

func (s *SchemaCounter[CV, CD]) SetTx(ctx context.Context, tx *gorm.DB, userSchema *model.UserSchema, value CV) error {
	err := gorm.
		G[model.UserSchema](tx).
		Exec(ctx, setSchemaCounterSQL,
			sql.Named("counter_name", s.name),
			sql.Named("counter_value", value),
			sql.Named("schema_id", userSchema.SchemaID),
		)

	if err != nil {
		return err
	}

	s.counterValue.Set(userSchema, value)
	return nil
}

func (s *SchemaCounter[CV, CD]) ChangeTx(ctx context.Context, tx *gorm.DB, userSchema *model.UserSchema, value CD) error {
	var updated []*updatedCounter[CV]

	err := gorm.
		G[model.UserSchema](tx).
		Raw(changeSchemaCounterSQL,
			sql.Named("counter_name", s.name),
			sql.Named("counter_delta", value),
			sql.Named("schema_id", userSchema.SchemaID),
		).
		Scan(ctx, &updated)

	if err != nil {
		return err
	}

	if len(updated) > 0 {
		s.counterValue.Set(userSchema, updated[0].Count)
	}

	return nil
}
