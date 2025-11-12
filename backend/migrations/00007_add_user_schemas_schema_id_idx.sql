-- +goose Up
CREATE INDEX IF NOT EXISTS idx_user_schemas_schema_id
  ON user_schemas (schema_id)
  INCLUDE (access, created_at);

-- +goose Down
DROP INDEX IF EXISTS idx_user_schemas_schema_id;
