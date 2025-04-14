import { defineStore } from 'pinia';
import { ref, toRef } from 'vue';
import type { ISchema } from '@/models';
import { type HttpBody, useHttpClient } from '@/composables';
import { useEditorSaveDispatcher } from '../composables';

type UpdateSchemaRequest = Partial<Omit<ISchema, 'id'>>;

export const useEditorStore = defineStore('schemas/editor', () => {
  const httpClient = useHttpClient();
  const schema = ref<ISchema>(null!);

  const saveDispatcher = useEditorSaveDispatcher(schema, async () => {
    await httpClient.patch<HttpBody, UpdateSchemaRequest>(['/schemas', schema.value.id], schema.value);
  });

  async function loadSchema(id: string): Promise<void> {
    schema.value = await httpClient.get(['/schemas', id]);
    saveDispatcher.enable();
  }

  function destroy() {
    saveDispatcher.disable();
    saveDispatcher.flush();
  }

  return {
    schema,
    loadSchema,
    destroy,
    hasUnsavedChanges: toRef(saveDispatcher, 'hasUnsavedChanges'),
  };
});
