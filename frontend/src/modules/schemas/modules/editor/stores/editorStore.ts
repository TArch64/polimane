import { defineStore } from 'pinia';
import { onScopeDispose, type Ref, ref, toRef } from 'vue';
import { useDebounceFn } from '@vueuse/core';
import {
  type ISchema,
  isRefBead,
  isSpannableBead,
  type SchemaBeads,
  type SchemaUpdate,
} from '@/models';
import { type HttpBody, HttpTransport, useAccessPermissions, useHttpClient } from '@/composables';
import { useSchemaBeadsLimit } from '@/composables/subscription';
import { getObjectEntries } from '@/helpers';
import { AccessLevel } from '@/enums';
import { useEditorSaveDispatcher } from './composables';
import { useHistoryStore } from './historyStore';

export const useEditorStore = defineStore('schemas/editor', () => {
  const http = useHttpClient();
  const schema: Ref<ISchema> = ref(null!);

  const historyStore = useHistoryStore();
  const permissions = useAccessPermissions(() => schema.value?.access ?? AccessLevel.READ);
  const schemaBeadsLimit = useSchemaBeadsLimit(schema);

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

  const onBeadsChange = useDebounceFn((beads: SchemaBeads) => {
    schemaBeadsLimit.current = Object
      .values(beads)
      .filter((bead) => !isRefBead(bead))
      .length;
  }, 100);

  const saveDispatcher = useEditorSaveDispatcher(schema, {
    async onSave(patch) {
      cleanupOrphanBeads(patch);

      await http.patch<HttpBody, SchemaUpdate>(['/schemas', schema.value.id], patch, {
        // Chrome has issues with fetch sending big request body
        transport: HttpTransport.LEGACY,
      });

      Object.assign(schema.value, {
        ...patch,
        updatedAt: new Date().toISOString(),
      });
    },

    onChange: { beads: onBeadsChange },
  });

  async function loadSchema(id: string): Promise<void> {
    schema.value = await http.get(['/schemas', id]);
    await historyStore.init(schema);
    saveDispatcher.enable();
  }

  async function deleteSchema(): Promise<void> {
    saveDispatcher.disable();
    saveDispatcher.abandon();
    await http.delete(['/schemas', schema.value.id]);
  }

  onScopeDispose(async () => {
    saveDispatcher.disable();
    await saveDispatcher.flush();
  });

  return {
    schema,
    loadSchema,
    deleteSchema,
    canEdit: toRef(permissions, 'write'),
    canEditAccess: toRef(permissions, 'admin'),
    canDelete: toRef(permissions, 'admin'),
    hasUnsavedChanges: toRef(saveDispatcher, 'hasUnsavedChanges'),
    isSaving: toRef(saveDispatcher, 'isSaving'),
    save: saveDispatcher.flush,
  };
});
