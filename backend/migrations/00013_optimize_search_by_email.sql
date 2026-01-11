-- +goose Up
DROP INDEX IF EXISTS idx_schema_invitations_schema_id;

CREATE INDEX IF NOT EXISTS idx_schema_invitations_schema_id
  ON schema_invitations (schema_id, email);

-- +goose Down
DROP INDEX IF EXISTS idx_schema_invitations_schema_id;

CREATE INDEX IF NOT EXISTS idx_schema_invitations_schema_id
  ON schema_invitations (schema_id);
