import { defineStore } from 'pinia';
import { computed } from 'vue';
import { useAsyncData, useHttpClient } from '@/composables';
import type { ISchema } from '@/models';

export interface ICreateSchemaInput {
  name: string;
}

export const useSchemasStore = defineStore('schemas', () => {
  const httpClient = useHttpClient();

  const schemas = useAsyncData({
    loader: () => httpClient.get<ISchema[]>('/schemas'),
    default: [],
  });

  const hasSchemas = computed(() => !!schemas.data.length);

  async function createSchema(input: ICreateSchemaInput): Promise<void> {
    const schema = await httpClient.post<ISchema, ICreateSchemaInput>('/schemas', input);
    schemas.data.push(schema);
  }

  return { schemas, hasSchemas, createSchema };
});
