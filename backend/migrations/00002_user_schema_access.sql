-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_schemas
  ADD COLUMN access smallint NOT NULL DEFAULT 3,
  ADD CONSTRAINT chk_user_schemas_access CHECK (access IN (1, 2, 3));

ALTER TABLE user_schemas
  ALTER COLUMN access DROP DEFAULT;

CREATE INDEX idx_user_schemas_users_access ON user_schemas (user_id, access);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_schemas
  DROP COLUMN access;
-- +goose StatementEnd
