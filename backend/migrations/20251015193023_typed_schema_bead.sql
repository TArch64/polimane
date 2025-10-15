UPDATE schemas
SET beads = (SELECT jsonb_object_agg(
                      key,
                      jsonb_build_object(
                        'color', value,
                        'kind', 'circle'
                      )
                    )
             FROM jsonb_each_text(beads));
