import { defineStore } from 'pinia';
import { computed } from 'vue';
import { useAsyncData, useHttpClient } from '@/composables';
import type { ISchema, ISchemaPattern } from '@/models';
import type { MarkOptional } from '@/types';

export type SchemaListItem = Omit<ISchema, 'content'>;

export interface ICreateSchemaInput {
  name: string;
  palette?: string[];
  content?: ISchemaPattern[];
}

type CreateSchemaRequest = MarkOptional<Omit<ISchema, 'id'>, 'palette' | 'content'>;

export const useSchemasStore = defineStore('schemas/list', () => {
  const httpClient = useHttpClient();

  const schemas = useAsyncData({
    loader: () => httpClient.get<SchemaListItem[]>('/schemas'),
    default: [],
  });

  const hasSchemas = computed(() => !!schemas.data.length);

  function createSchema(input: ICreateSchemaInput): Promise<SchemaListItem> {
    return httpClient.post<SchemaListItem, CreateSchemaRequest>('/schemas', input);
  }

  async function deleteSchema(deletingSchema: ISchema): Promise<void> {
    await httpClient.delete(['/schemas', deletingSchema.id]);
    schemas.data = schemas.data.filter((schema) => schema.id !== deletingSchema.id);
  }

  async function copySchema(copyingSchema: ISchema): Promise<SchemaListItem> {
    return httpClient.post(['/schemas', copyingSchema.id, 'copy'], {});
  }

  return { schemas, hasSchemas, createSchema, deleteSchema, copySchema };
});
