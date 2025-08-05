-- Modify "schemas" table
ALTER TABLE "public"."schemas"
    ADD COLUMN "screenshoted_at" timestamptz NULL;
