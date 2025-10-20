UPDATE schemas
SET beads = CASE
              WHEN beads = '{}'::jsonb THEN '{}'::jsonb
              ELSE (SELECT jsonb_object_agg(
                             key,
                             jsonb_build_object(
                               'kind', 'circle',
                               'circle', jsonb_build_object('color', value)
                             )
                           )
                    FROM jsonb_each_text(schemas.beads))
  END;

