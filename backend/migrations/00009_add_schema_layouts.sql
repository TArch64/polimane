-- +goose Up
CREATE TYPE schema_layout AS enum ('linear', 'radial');

ALTER TABLE schemas
  ADD COLUMN IF NOT EXISTS layout schema_layout NOT NULL DEFAULT 'linear';

-- +goose Down
ALTER TABLE schemas
  DROP COLUMN IF EXISTS layout;

DROP TYPE IF EXISTS schema_layout;
