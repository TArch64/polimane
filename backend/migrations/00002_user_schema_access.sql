-- +goose Up
-- +goose StatementBegin
CREATE TYPE access_level AS enum ('read', 'write', 'admin');

ALTER TABLE user_schemas
  ADD COLUMN access access_level NOT NULL DEFAULT 'admin';

ALTER TABLE user_schemas
  ALTER COLUMN access DROP DEFAULT;

CREATE INDEX idx_user_schemas_users_access ON user_schemas (user_id, access);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_schemas
  DROP COLUMN access;

DROP INDEX idx_user_schemas_users_access;
DROP TYPE access_level;
-- +goose StatementEnd
