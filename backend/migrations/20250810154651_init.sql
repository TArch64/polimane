-- Create "schemas" table
CREATE TABLE "public"."schemas"
(
    "id"              uuid                   NOT NULL DEFAULT gen_random_uuid(),
    "created_at"      timestamptz            NOT NULL,
    "updated_at"      timestamptz            NOT NULL,
    "name"            character varying(255) NOT NULL,
    "palette"         jsonb                  NOT NULL,
    "size"            jsonb                  NOT NULL,
    "beads"           jsonb                  NOT NULL,
    "screenshoted_at" timestamptz            NULL,
    PRIMARY KEY ("id")
);
-- Create index "idx_schemas_name" to table: "schemas"
CREATE INDEX "idx_schemas_name" ON "public"."schemas" ("name");
-- Create "users" table
CREATE TABLE "public"."users"
(
    "id"        uuid                  NOT NULL DEFAULT gen_random_uuid(),
    "workos_id" character varying(32) NOT NULL,
    PRIMARY KEY ("id")
);
-- Create index "idx_users_workos_id" to table: "users"
CREATE UNIQUE INDEX "idx_users_workos_id" ON "public"."users" ("workos_id");
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
