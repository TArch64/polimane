import { defineStore } from 'pinia';
import { computed, ref, toRef } from 'vue';
import type { ISchema, ISchemaPattern } from '@/models';
import { type HttpBody, useHttpClient } from '@/composables';
import { useEditorSaveDispatcher } from '../composables';

type UpdateSchemaRequest = Partial<Omit<ISchema, 'id'>>;

export const useEditorStore = defineStore('schemas/editor', () => {
  const httpClient = useHttpClient();
  const schema = ref<ISchema>(null!);

  const saveDispatcher = useEditorSaveDispatcher(schema, async (patch) => {
    await httpClient.patch<HttpBody, UpdateSchemaRequest>(['/schemas', schema.value.id], patch);
  });

  const activePatternId = ref<string | null>(null);

  const activePattern = computed(() => {
    return activePatternId.value
      ? schema.value.content.patterns.find((pattern) => pattern.id === activePatternId.value)
      : null;
  });

  function activatePattern(pattern: ISchemaPattern): void {
    activePatternId.value = pattern.id;
  }

  async function loadSchema(id: string): Promise<void> {
    schema.value = await httpClient.get(['/schemas', id]);
    activePatternId.value = schema.value.content.patterns[0]?.id ?? null;
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
    activePattern,
    activatePattern,
    hasUnsavedChanges: toRef(saveDispatcher, 'hasUnsavedChanges'),
  };
});
