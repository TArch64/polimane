import { defineStore } from 'pinia';
import { computed, onScopeDispose, type Ref, ref, toRef } from 'vue';
import type { ISchema } from '@/models';
import { type HttpBody, useAccessPermissions, useHttpClient } from '@/composables';
import { AccessLevel } from '@/enums';
import { useSchemasCreatedCounter } from '@/composables/subscription';
import {
  useEditorBeadsLimitUpdater,
  useEditorSaveDispatcher,
  useEditorSaveProcessor,
} from './composables';
import { useHistoryStore } from './historyStore';

export interface ISchemasRequest {
  ids: string[];
}

export const useEditorStore = defineStore('schemas/editor', () => {
  const http = useHttpClient();
  const schema: Ref<ISchema> = ref(null!);

  const historyStore = useHistoryStore();

  const permissions = useAccessPermissions(() => schema.value?.access ?? AccessLevel.READ);
  const saveEditor = useEditorSaveProcessor(schema);
  const beadsLimitUpdater = useEditorBeadsLimitUpdater(schema);
  const schemasCreatedCounter = useSchemasCreatedCounter();

  const saveDispatcher = useEditorSaveDispatcher(schema, {
    onSave: saveEditor,
    onChange: { beads: beadsLimitUpdater.onBeadsChange },
  });

  async function loadSchema(id: string): Promise<void> {
    schema.value = await http.get(['/schemas', id]);
    await historyStore.init(schema);
    saveDispatcher.enable();
  }

  async function deleteSchema(): Promise<void> {
    try {
      saveDispatcher.disable();

      await http.delete<HttpBody, ISchemasRequest>(['/schemas', 'delete'], {
        ids: [schema.value.id],
      });

      saveDispatcher.abandon();
    } catch (error) {
      saveDispatcher.enable();
      throw error;
    }
  }

  const canEdit = computed(() => {
    return permissions.write && !schemasCreatedCounter.isReached;
  });

  onScopeDispose(async () => {
    saveDispatcher.disable();
    await saveDispatcher.flush();
    beadsLimitUpdater.destroy();
  });

  return {
    schema,
    loadSchema,
    deleteSchema,
    canEdit,
    canEditName: toRef(permissions, 'write'),
    canEditAccess: toRef(permissions, 'admin'),
    canDelete: toRef(permissions, 'admin'),
    hasUnsavedChanges: toRef(saveDispatcher, 'hasUnsavedChanges'),
    isSaving: toRef(saveDispatcher, 'isSaving'),
    save: saveDispatcher.flush,
  };
});
