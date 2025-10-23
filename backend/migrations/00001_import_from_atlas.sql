-- +goose Up
-- +goose StatementBegin
CREATE TABLE schemas
(
  id               uuid                   NOT NULL DEFAULT gen_random_uuid(),
  created_at       timestamptz            NOT NULL,
  updated_at       timestamptz            NOT NULL,
  name             character varying(255) NOT NULL,
  palette          jsonb                  NOT NULL,
  size             jsonb                  NOT NULL,
  beads            jsonb                  NOT NULL,
  screenshoted_at  timestamptz            NULL,
  background_color character varying(30)  NOT NULL DEFAULT '#f8f8f8',
  PRIMARY KEY (id)
);

CREATE INDEX idx_schemas_name ON schemas (name);

CREATE TABLE users
(
  id        uuid                  NOT NULL DEFAULT gen_random_uuid(),
  workos_id character varying(32) NOT NULL,
  PRIMARY KEY (id)
);

CREATE UNIQUE INDEX idx_users_workos_id ON users (workos_id);

CREATE TABLE user_schemas
(
  created_at timestamptz NOT NULL,
  updated_at timestamptz NOT NULL,
  user_id    uuid        NOT NULL,
  schema_id  uuid        NOT NULL,
  PRIMARY KEY (user_id, schema_id),
  CONSTRAINT fk_user_schemas_schema FOREIGN KEY (schema_id) REFERENCES schemas (id) ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT fk_user_schemas_user FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE NO ACTION ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_schemas;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS schemas;
-- +goose StatementEnd
