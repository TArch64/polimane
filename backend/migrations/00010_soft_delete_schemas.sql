-- +goose Up
ALTER TABLE schemas
  ADD COLUMN IF NOT EXISTS deleted_at timestamptz NULL;

CREATE INDEX IF NOT EXISTS idx_schemas_deleted_at
  ON schemas (deleted_at)
  WHERE deleted_at IS NOT NULL;

-- +goose Down
ALTER TABLE schemas
  DROP COLUMN IF EXISTS deleted_at;
