-- +goose Up
ALTER TABLE users
  ADD COLUMN email      varchar(255) UNIQUE,
  ADD COLUMN first_name varchar(255),
  ADD COLUMN last_name  varchar(255);

-- +goose Down
ALTER TABLE users
  DROP COLUMN email,
  DROP COLUMN first_name,
  DROP COLUMN last_name;
