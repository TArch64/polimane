-- +goose Up
CREATE TABLE IF NOT EXISTS folders
(
  id         uuid                   NOT NULL DEFAULT GEN_RANDOM_UUID(),
  created_at timestamptz            NOT NULL,
  updated_at timestamptz            NOT NULL,
  name       character varying(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_folders_created_at
  ON folders (created_at);

CREATE TABLE IF NOT EXISTS folder_schemas
(
  folder_id  uuid        NOT NULL,
  schema_id  uuid        NOT NULL,
  created_at timestamptz NOT NULL,
  updated_at timestamptz NOT NULL,
  PRIMARY KEY (folder_id, schema_id),

  CONSTRAINT fk_folder_schemas_folder
    FOREIGN KEY (folder_id)
      REFERENCES folders (id)
      ON UPDATE NO ACTION
      ON DELETE CASCADE,

  CONSTRAINT fk_folder_schemas_schema
    FOREIGN KEY (schema_id)
      REFERENCES schemas (id)
      ON UPDATE NO ACTION
      ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_schemas_created_at
  ON schemas (created_at);

-- +goose Down
DROP TABLE IF EXISTS folder_schemas;
DROP TABLE IF EXISTS folders;
DROP INDEX IF EXISTS idx_schemas_created_at;
