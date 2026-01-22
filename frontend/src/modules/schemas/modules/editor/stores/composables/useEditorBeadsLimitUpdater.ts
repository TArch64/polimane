import type { Ref } from 'vue';
import { useDebounceFn } from '@vueuse/core';
import { type ISchema, isRefBead, type SchemaBeads } from '@/models';
import { useSchemaBeadsLimit } from '@/composables/subscription';
import { UpdateCountersMiddleware, useHttpClient } from '@/composables';

export interface IEditorBeadsLimitUpdater {
  onBeadsChange: (beads: SchemaBeads) => void;
  destroy: () => void;
}

export function useEditorBeadsLimitUpdater(schema: Ref<ISchema>): IEditorBeadsLimitUpdater {
  const http = useHttpClient();
  const middleware = http.getMiddleware(UpdateCountersMiddleware)!;

  const limit = useSchemaBeadsLimit(schema);

  const onBeadsChange = useDebounceFn((beads: SchemaBeads) => {
    limit.current = Object
      .values(beads)
      .filter((bead) => !isRefBead(bead))
      .length;
  }, 200);

  const destroy = middleware.onSchemaUpdate.listen((schemaId, counters) => {
    if (schemaId === schema.value.id) {
      schema.value.counters = counters;
    }
  });

  return {
    onBeadsChange,
    destroy,
  };
}
