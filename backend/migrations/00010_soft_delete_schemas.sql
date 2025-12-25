-- +goose Up
ALTER TABLE schemas
  ADD COLUMN IF NOT EXISTS deleted_at timestamptz NULL;

CREATE INDEX IF NOT EXISTS idx_schemas_deleted_at
  ON schemas (deleted_at)
  WHERE deleted_at IS NOT NULL;

ALTER TABLE user_schemas
  ADD COLUMN IF NOT EXISTS deleted_at timestamptz NULL;

CREATE INDEX IF NOT EXISTS idx_user_schemas_deleted_at
  ON user_schemas (user_id, deleted_at)
  WHERE deleted_at IS NOT NULL;


-- +goose Down
DROP INDEX
  IF EXISTS idx_schemas_deleted_at;

ALTER TABLE schemas
  DROP COLUMN IF EXISTS deleted_at;

DROP INDEX
  IF EXISTS idx_user_schemas_deleted_at;

ALTER TABLE user_schemas
  DROP COLUMN IF EXISTS deleted_at;

