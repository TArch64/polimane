-- +goose Up
CREATE TABLE schema_invitations
(
  email      varchar(255) NOT NULL,
  schema_id  uuid         NOT NULL,
  access     smallint     NOT NULL,
  expires_at timestamptz  NOT NULL,
  PRIMARY KEY (email, schema_id),

  CONSTRAINT fk_schema_invitations_schema
    FOREIGN KEY (schema_id)
      REFERENCES schemas (id)
      ON DELETE CASCADE,

  CONSTRAINT chk_schema_invitations_access
    CHECK (access IN (1, 2, 3))
);

CREATE INDEX idx_schema_invitations_schema_id
  ON schema_invitations (schema_id);

CREATE INDEX idx_schema_invitations_expires_at
  ON schema_invitations (expires_at);

-- +goose Down
DROP TABLE schema_invitations;
