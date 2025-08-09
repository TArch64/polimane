import { defineStore } from 'pinia';
import { onScopeDispose, type Ref, ref, toRef } from 'vue';
import type { ISchema } from '@/models';
import { type HttpBody, useHttpClient } from '@/composables';
import { useEditorHistory, useEditorSaveDispatcher } from '../composables';

type UpdateSchemaRequest = Partial<Omit<ISchema, 'id'>>;

interface IUpdateScreenshotRequest {
  src: string;
}

export const useEditorStore = defineStore('schemas/editor', () => {
  const http = useHttpClient();
  const schema: Ref<ISchema> = ref(null!);

  const history = useEditorHistory(schema);

  const saveDispatcher = useEditorSaveDispatcher(schema, async (patch) => {
    await http.patch<HttpBody, UpdateSchemaRequest>(['/schemas', schema.value.id], patch);
    schema.value.updatedAt = new Date().toISOString();
  });

  async function loadSchema(id: string): Promise<void> {
    schema.value = await http.get(['/schemas', id]);
    await history.init();
    saveDispatcher.enable();
  }

  async function deleteSchema(): Promise<void> {
    saveDispatcher.disable();
    saveDispatcher.abandon();
    history.destroy();
    await http.delete(['/schemas', schema.value.id]);
  }

  async function updateScreenshot(src: string): Promise<void> {
    await http.patch<HttpBody, IUpdateScreenshotRequest>(['/schemas', schema.value.id, 'screenshot'], { src });
    schema.value.screenshotedAt = new Date().toISOString();
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
    updateScreenshot,
    hasUnsavedChanges: toRef(saveDispatcher, 'hasUnsavedChanges'),
    isSaving: toRef(saveDispatcher, 'isSaving'),
    save: saveDispatcher.flush,
    onSaved: saveDispatcher.onSaved,
    undo: history.undo,
    canUndo: toRef(history, 'canUndo'),
    redo: history.redo,
    canRedo: toRef(history, 'canRedo'),
  };
});
