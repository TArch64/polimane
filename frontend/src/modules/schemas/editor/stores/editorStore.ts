import { defineStore } from 'pinia';
import { onScopeDispose, type Ref, ref, toRef } from 'vue';
import type { ISchema } from '@/models';
import { type HttpBody, HttpTransport, useHttpClient } from '@/composables';
import { useEditorSaveDispatcher } from '../composables';
import { useHistoryStore } from './historyStore';

type UpdateSchemaRequest = Partial<Omit<ISchema, 'id'>>;

export const useEditorStore = defineStore('schemas/editor', () => {
  const http = useHttpClient();
  const schema: Ref<ISchema> = ref(null!);

  const historyStore = useHistoryStore();

  const saveDispatcher = useEditorSaveDispatcher(schema, async (patch) => {
    await http.patch<HttpBody, UpdateSchemaRequest>(['/schemas', schema.value.id], patch, {
      // Chrome has issues with fetch sending big request body
      transport: HttpTransport.LEGACY,
    });

    schema.value.updatedAt = new Date().toISOString();
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
    hasUnsavedChanges: toRef(saveDispatcher, 'hasUnsavedChanges'),
    isSaving: toRef(saveDispatcher, 'isSaving'),
    save: saveDispatcher.flush,
  };
});
