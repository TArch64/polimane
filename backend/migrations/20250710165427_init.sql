-- Create "schemas" table
CREATE TABLE "public"."schemas"
(
    "id"         uuid                   NOT NULL DEFAULT gen_random_uuid(),
    "created_at" timestamptz            NOT NULL,
    "updated_at" timestamptz            NOT NULL,
    "name"       character varying(255) NOT NULL,
    "palette"    jsonb                  NOT NULL,
    "content"    jsonb                  NOT NULL,
    PRIMARY KEY ("id")
);
-- Create index "idx_schemas_name" to table: "schemas"
CREATE INDEX "idx_schemas_name" ON "public"."schemas" ("name");
-- Create "users" table
CREATE TABLE "public"."users"
(
    "id"            uuid                   NOT NULL DEFAULT gen_random_uuid(),
    "created_at"    timestamptz            NOT NULL,
    "updated_at"    timestamptz            NOT NULL,
    "name"          character varying(255) NOT NULL,
    "password_hash" text                   NOT NULL,
    PRIMARY KEY ("id")
);
-- Create index "idx_users_name" to table: "users"
CREATE UNIQUE INDEX "idx_users_name" ON "public"."users" ("name");
-- Create "user_schemas" table
CREATE TABLE "public"."user_schemas"
(
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "user_id"    uuid        NOT NULL,
    "schema_id"  uuid        NOT NULL,
    PRIMARY KEY ("user_id", "schema_id"),
    CONSTRAINT "fk_user_schemas_schema" FOREIGN KEY ("schema_id") REFERENCES "public"."schemas" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
    CONSTRAINT "fk_user_schemas_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
