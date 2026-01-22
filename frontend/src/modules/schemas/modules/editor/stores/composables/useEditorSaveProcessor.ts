import type { Ref } from 'vue';
import { type ISchema, isRefBead, isSpannableBead, type SchemaUpdate } from '@/models';
import { getObjectEntries } from '@/helpers';
import { type HttpBody, HttpTransport, useHttpClient } from '@/composables';

export function useEditorSaveProcessor(schema: Ref<ISchema>) {
  const http = useHttpClient();

  function cleanupOrphanBeads(patch: Partial<ISchema>): void {
    if (!patch.beads) {
      return;
    }

    for (const [coord, bead] of getObjectEntries(patch.beads)) {
      if (isRefBead(bead)) {
        const targetBead = patch.beads[bead.ref!.to];

        if (!targetBead || !isSpannableBead(targetBead)) {
          delete patch.beads[coord];
        }
      }
    }
  }

  return async (patch: SchemaUpdate) => {
    cleanupOrphanBeads(patch);

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
