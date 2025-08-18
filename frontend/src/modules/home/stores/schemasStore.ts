import { defineStore } from 'pinia';
import { computed, nextTick } from 'vue';
import { useAsyncData, useHttpClient, useRouteTransition } from '@/composables';
import type { ISchema } from '@/models';

export type SchemaListItem = Omit<ISchema, 'beads' | 'size'>;

export interface ICreateSchemaInput {
  name: string;
  palette?: string[];
}

export const useSchemasStore = defineStore('schemas/list', () => {
  const routeTransition = useRouteTransition();
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
    routeTransition.start(() => {
      schemas.makeOptimisticUpdate((schemas) => {
        return schemas.filter((schema) => schema.id !== deletingSchema.id);
      });
      return nextTick();
    });

    try {
      await http.delete(['/schemas', deletingSchema.id]);
      schemas.commitOptimisticUpdate();
    } catch (error) {
      routeTransition.start(() => {
        schemas.rollbackOptimisticUpdate();
        return nextTick();
      });
      throw error;
    }
  }

  async function copySchema(copyingSchema: ISchema): Promise<SchemaListItem> {
    return http.post(['/schemas', copyingSchema.id, 'copy'], {});
  }

  return { schemas, hasSchemas, createSchema, deleteSchema, copySchema };
});
