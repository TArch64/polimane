-- +goose Up
ALTER TABLE schemas
  ADD COLUMN IF NOT EXISTS deleted_by uuid NULL;

CREATE INDEX IF NOT EXISTS
  idx_schemas_deleted_by ON schemas (deleted_by)
  WHERE deleted_by IS NOT NULL;

-- +goose Down
DROP INDEX IF EXISTS idx_schemas_deleted_by;

ALTER TABLE schemas
  DROP COLUMN IF EXISTS deleted_by;
