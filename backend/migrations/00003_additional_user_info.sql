-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
  ADD COLUMN email      varchar(255) UNIQUE,
  ADD COLUMN first_name varchar(255),
  ADD COLUMN last_name  varchar(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
  DROP COLUMN email,
  DROP COLUMN first_name,
  DROP COLUMN last_name;
-- +goose StatementEnd
