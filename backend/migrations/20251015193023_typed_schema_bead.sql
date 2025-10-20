UPDATE schemas
SET beads = (SELECT jsonb_object_agg(
                      key,
                      jsonb_build_object(
                        'kind', 'circle',
                        'circle', jsonb_build_object('color', value)
                      )
                    )
             FROM jsonb_each_text(schemas.beads));
