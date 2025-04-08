import { defineStore } from 'pinia';
import { ref } from 'vue';
import type { ISchema } from '@/models';
import { useHttpClient } from '@/composables';

export const useEditorStore = defineStore('schemas/editor', () => {
  const httpClient = useHttpClient();
  const schema = ref<ISchema>(null!);

  async function loadSchema(id: string): Promise<void> {
    schema.value = await httpClient.get(['/schemas', id]);
  }

  return { schema, loadSchema };
});
