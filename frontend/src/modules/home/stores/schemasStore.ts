import { defineStore } from 'pinia';
import { computed } from 'vue';
import { useAsyncData, useHttpClient } from '@/composables';
import type { ISchema, ISchemaPattern } from '@/models';

export type SchemaListItem = Omit<ISchema, 'content'>;

export interface ICreateSchemaInput {
  name: string;
  palette?: string[];
  content?: ISchemaPattern[];
}

export const useSchemasStore = defineStore('schemas/list', () => {
  const http = useHttpClient();

  const schemas = useAsyncData({
    loader: () => http.get<SchemaListItem[]>('/schemas'),
    default: [],
  });

  const hasSchemas = computed(() => !!schemas.data.length);

  function createSchema(input: ICreateSchemaInput): Promise<SchemaListItem> {
    return http.post<SchemaListItem, ICreateSchemaInput>('/schemas', input);
  }

  async function deleteSchema(deletingSchema: ISchema): Promise<void> {
    await http.delete(['/schemas', deletingSchema.id]);
    schemas.data = schemas.data.filter((schema) => schema.id !== deletingSchema.id);
  }

  async function copySchema(copyingSchema: ISchema): Promise<SchemaListItem> {
    return http.post(['/schemas', copyingSchema.id, 'copy'], {});
  }

  return { schemas, hasSchemas, createSchema, deleteSchema, copySchema };
});
