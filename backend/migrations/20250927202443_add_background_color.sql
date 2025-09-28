-- Modify "schemas" table
ALTER TABLE "public"."schemas"
    ADD COLUMN "background_color" character varying(30) NOT NULL DEFAULT '#f8f8f8';
