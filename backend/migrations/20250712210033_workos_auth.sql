-- Clear Invalid Data
DELETE
FROM users;
DELETE
FROM schemas;
-- Modify "users" table
ALTER TABLE "public"."users"
    DROP COLUMN "created_at",
    DROP COLUMN "updated_at",
    DROP COLUMN "name",
    DROP COLUMN "password_hash",
    ADD COLUMN "workos_id" character varying(32) NOT NULL;
-- Create index "idx_users_workos_id" to table: "users"
CREATE UNIQUE INDEX "idx_users_workos_id" ON "public"."users" ("workos_id");
