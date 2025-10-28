-- +goose Up
UPDATE schemas
SET name = REGEXP_REPLACE(name, ' +', ' ', 'g');
