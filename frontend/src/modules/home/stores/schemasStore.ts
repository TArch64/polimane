import { defineStore } from 'pinia';
import { computed } from 'vue';
import { useAsyncData, useHttpClient } from '@/composables';
import type { ISchema } from '@/models';

export type SchemaListItem = Omit<ISchema, 'content'>;

export interface ICreateSchemaInput {
  name: string;
}

type CreateSchemaRequest = Omit<ISchema, 'id' | 'palette'>;

export const useSchemasStore = defineStore('schemas/list', () => {
  const httpClient = useHttpClient();

  const schemas = useAsyncData({
    loader: () => httpClient.get<SchemaListItem[]>('/schemas'),
    default: [],
  });

  const hasSchemas = computed(() => !!schemas.data.length);

  function createSchema(input: ICreateSchemaInput): Promise<SchemaListItem> {
    return httpClient.post<SchemaListItem, CreateSchemaRequest>('/schemas', {
      ...input,
      content: [],
    });
  }

  async function deleteSchema(deletingSchema: ISchema): Promise<void> {
    await httpClient.delete(['/schemas', deletingSchema.id]);
    schemas.data = schemas.data.filter((schema) => schema.id !== deletingSchema.id);
  }

  return { schemas, hasSchemas, createSchema, deleteSchema };
});
