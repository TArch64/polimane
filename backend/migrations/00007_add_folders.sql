-- +goose Up
CREATE TABLE IF NOT EXISTS folders
(
  id         uuid                   NOT NULL DEFAULT GEN_RANDOM_UUID(),
  created_at timestamptz            NOT NULL,
  updated_at timestamptz            NOT NULL,
  name       character varying(255) NOT NULL,
  user_id    uuid                   NOT NULL,
  PRIMARY KEY (id),

  CONSTRAINT fk_folders_user
    FOREIGN KEY (user_id)
      REFERENCES users (id)
      ON UPDATE NO ACTION
      ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_folders_created_at
  ON folders (created_at);

CREATE INDEX IF NOT EXISTS idx_folders_user_id
  ON folders (user_id);

ALTER TABLE user_schemas
  ADD COLUMN IF NOT EXISTS folder_id uuid NULL,
  ADD CONSTRAINT fk_user_schemas_folder
    FOREIGN KEY (folder_id)
      REFERENCES folders (id)
      ON UPDATE NO ACTION
      ON DELETE SET NULL;

CREATE INDEX IF NOT EXISTS idx_user_schemas_folder_id
  ON user_schemas (folder_id);

-- +goose Down
DROP TABLE IF EXISTS folders;

ALTER TABLE user_schemas
  DROP CONSTRAINT IF EXISTS fk_user_schemas_folder,
  DROP COLUMN IF EXISTS folder_id;
