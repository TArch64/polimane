import { defineStore } from 'pinia';
import { onScopeDispose, type Ref, ref, toRef } from 'vue';
import type { ISchema } from '@/models';
import { useAccessPermissions, useHttpClient } from '@/composables';
import { AccessLevel } from '@/enums';
import {
  useEditorBeadsLimitUpdater,
  useEditorSaveDispatcher,
  useEditorSaveProcessor,
} from './composables';
import { useHistoryStore } from './historyStore';

export const useEditorStore = defineStore('schemas/editor', () => {
  const http = useHttpClient();
  const schema: Ref<ISchema> = ref(null!);

  const historyStore = useHistoryStore();

  const permissions = useAccessPermissions(() => schema.value?.access ?? AccessLevel.READ);
  const saveEditor = useEditorSaveProcessor(schema);
  const beadsLimitUpdater = useEditorBeadsLimitUpdater(schema);

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
    saveDispatcher.disable();
    saveDispatcher.abandon();
    await http.delete(['/schemas', schema.value.id]);
  }

  onScopeDispose(async () => {
    saveDispatcher.disable();
    await saveDispatcher.flush();
    beadsLimitUpdater.destroy();
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
