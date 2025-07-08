import { defineStore } from 'pinia';
import { onScopeDispose, type Ref, ref, toRef } from 'vue';
import type { ISchema } from '@/models';
import { type HttpBody, useHttpClient } from '@/composables';
import { useEditorHistory, useEditorSaveDispatcher } from '../composables';
import { setSchemaRelations } from '../models';

type UpdateSchemaRequest = Partial<Omit<ISchema, 'id'>>;

export const useEditorStore = defineStore('schemas/editor', () => {
  const httpClient = useHttpClient();
  const schema: Ref<ISchema> = ref(null!);

  const history = useEditorHistory(schema);

  const saveDispatcher = useEditorSaveDispatcher(schema, async (patch) => {
    await httpClient.patch<HttpBody, UpdateSchemaRequest>(['/schemas', schema.value.id], patch);
  });

  async function loadSchema(id: number): Promise<void> {
    schema.value = await httpClient.get(['/schemas', id]);
    schema.value.content ??= [];
    setSchemaRelations(schema.value);

    await history.init();
    saveDispatcher.enable();
  }

  async function deleteSchema(): Promise<void> {
    saveDispatcher.disable();
    saveDispatcher.abandon();
    history.destroy();
    await httpClient.delete(['/schemas', schema.value.id]);
  }

  onScopeDispose(async () => {
    saveDispatcher.disable();
    history.destroy();
    await saveDispatcher.flush();
  });

  return {
    schema,
    loadSchema,
    deleteSchema,
    hasUnsavedChanges: toRef(saveDispatcher, 'hasUnsavedChanges'),
    undo: history.undo,
    canUndo: toRef(history, 'canUndo'),
    redo: history.redo,
    canRedo: toRef(history, 'canRedo'),
  };
});
