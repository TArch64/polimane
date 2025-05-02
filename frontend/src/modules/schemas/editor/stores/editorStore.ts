import { defineStore } from 'pinia';
import { ref, toRef } from 'vue';
import type { ISchema } from '@/models';
import { type HttpBody, useHttpClient } from '@/composables';
import { useEditorSaveDispatcher } from '../composables';
import { setObjectParent } from '../models';

type UpdateSchemaRequest = Partial<Omit<ISchema, 'id'>>;

export const useEditorStore = defineStore('schemas/editor', () => {
  const httpClient = useHttpClient();
  const schema = ref<ISchema>(null!);

  const saveDispatcher = useEditorSaveDispatcher(schema, async (patch) => {
    await httpClient.patch<HttpBody, UpdateSchemaRequest>(['/schemas', schema.value.id], patch);
  });

  async function loadSchema(id: string): Promise<void> {
    schema.value = await httpClient.get(['/schemas', id]);
    schema.value.content ??= [];

    for (const object of schema.value.content) {
      setObjectParent(schema.value, object);
    }

    saveDispatcher.enable();
  }

  async function destroy() {
    saveDispatcher.disable();
    await saveDispatcher.flush();
  }

  async function deleteSchema(): Promise<void> {
    saveDispatcher.disable();
    saveDispatcher.abandon();
    await httpClient.delete(['/schemas', schema.value.id]);
  }

  return {
    schema,
    destroy,
    loadSchema,
    deleteSchema,
    hasUnsavedChanges: toRef(saveDispatcher, 'hasUnsavedChanges'),
  };
});
