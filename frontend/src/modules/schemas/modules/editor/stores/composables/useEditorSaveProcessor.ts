import type { Ref } from 'vue';
import {
  type BeadCoord,
  type ISchema,
  isRefBead,
  isSpannableBead,
  type SchemaUpdate,
} from '@/models';
import { deleteObjectKeys, getObjectEntries } from '@/helpers';
import { type HttpBody, HttpTransport, useHttpClient } from '@/composables';
import { useSchemaBeadsCounter } from '@/composables/subscription';

export function useEditorSaveProcessor(schema: Ref<ISchema>) {
  const http = useHttpClient();
  const schemaBeadsCounter = useSchemaBeadsCounter(schema);

  function cleanupOrphanRefs(patch: Partial<ISchema>): void {
    if (!patch.beads) {
      return;
    }

    const removingSet = new Set<BeadCoord>();

    for (const [coord, bead] of getObjectEntries(patch.beads)) {
      if (isRefBead(bead)) {
        const targetBead = patch.beads[bead.ref!.to];

        if (!targetBead || !isSpannableBead(targetBead)) {
          removingSet.add(coord);
        }
      }
    }

    if (removingSet.size > 0) {
      patch.beads = deleteObjectKeys(patch.beads, removingSet);
    }
  }

  return async (patch: SchemaUpdate) => {
    cleanupOrphanRefs(patch);

    if ('beads' in patch && schemaBeadsCounter.max && schemaBeadsCounter.current > schemaBeadsCounter.max) {
      return;
    }

    await http.patch<HttpBody, SchemaUpdate>(['/schemas', schema.value.id], patch, {
      // Chrome has issues with fetch sending big request body
      transport: HttpTransport.LEGACY,
    });

    Object.assign(schema.value, {
      ...patch,
      updatedAt: new Date().toISOString(),
    });
  };
}
