import type { Ref } from 'vue';
import { useThrottleFn } from '@vueuse/core';
import { type ISchema, isRefBead, type SchemaBeads } from '@/models';
import { useSchemaBeadsCounter } from '@/composables/subscription';
import { UpdateCountersMiddleware, useHttpClient } from '@/composables';

export interface IEditorBeadsLimitUpdater {
  onBeadsChange: (beads: SchemaBeads) => void;
  destroy: () => void;
}

export function useEditorBeadsLimitUpdater(schema: Ref<ISchema>): IEditorBeadsLimitUpdater {
  const http = useHttpClient();
  const middleware = http.getMiddleware(UpdateCountersMiddleware)!;

  const counter = useSchemaBeadsCounter(schema);

  const onBeadsChange = useThrottleFn((beads: SchemaBeads) => {
    counter.current = Object
      .values(beads)
      .filter((bead) => !isRefBead(bead))
      .length;
  }, 200, true);

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
