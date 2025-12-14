-- +goose Up
ALTER TABLE users
  ADD COLUMN IF NOT EXISTS deleted_at timestamptz NULL;

CREATE INDEX IF NOT EXISTS idx_users_deleted_at
  ON users (deleted_at);

-- +goose Down
ALTER TABLE users
  DROP COLUMN IF EXISTS deleted_at;
