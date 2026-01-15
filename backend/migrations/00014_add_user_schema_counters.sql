-- +goose Up
ALTER TABLE user_schemas
  ADD COLUMN IF NOT EXISTS counters jsonb DEFAULT '{}'::jsonb;


WITH schema_users_count AS (SELECT schema_id AS id, COUNT(user_id)
                            FROM user_schemas
                            GROUP BY schema_id),

     schema_beads_count AS (SELECT id, COUNT(*)
                            FROM schemas,
                                 LATERAL JSONB_EACH(beads) AS schema_beads
                            WHERE schema_beads.value ->> 'kind' != 'ref'
                            GROUP BY id)
UPDATE user_schemas
SET counters = JSON_BUILD_OBJECT(
  'sharedAccess', schema_users_count.count,
  'schemaBeads', COALESCE(schema_beads_count.count, 0)
               )
FROM schema_users_count
       LEFT JOIN schema_beads_count ON schema_beads_count.id = schema_users_count.id
WHERE user_schemas.schema_id = schema_users_count.id;


CREATE
OR
REPLACE
FUNCTION jsonb_increment(
    target jsonb,
    KEY TEXT,
    delta SMALLINT
) RETURNS jsonb STABLE LANGUAGE SQL AS $$
    SELECT jsonb_set(
        target,
        ARRAY[key],
        TO_JSONB(COALESCE((target ->> key)::smallint, 0) + delta)
    )
$$;

-- +goose Down
DROP
FUNCTION jsonb_increment;

ALTER TABLE user_schemas
  DROP COLUMN IF EXISTS counters;
