import { defineStore } from 'pinia';
import { computed, nextTick } from 'vue';
import { useAsyncData, useHttpClient, useRouteTransition } from '@/composables';
import type { ISchema } from '@/models';

export type SchemaListItem = Omit<ISchema, 'beads' | 'size'>;

interface IListResponse {
  list: SchemaListItem[];
  total: number;
}

export interface ICreateSchemaInput {
  name: string;
  palette?: string[];
}

export const useSchemasStore = defineStore('schemas/list', () => {
  const routeTransition = useRouteTransition();
  const http = useHttpClient();

  const list = useAsyncData({
    loader: () => http.get<IListResponse>('/schemas'),
    default: { list: [], total: 0 },
  });

  const schemas = computed(() => list.data.list);
  const hasSchemas = computed(() => !!schemas.value.length);

  function createSchema(input: ICreateSchemaInput): Promise<SchemaListItem> {
    return http.post<SchemaListItem, ICreateSchemaInput>('/schemas', input);
  }

  async function deleteSchema(deletingSchema: ISchema): Promise<void> {
    routeTransition.start(() => {
      list.makeOptimisticUpdate(({ list, total }) => ({
        list: list.filter((schema) => schema.id !== deletingSchema.id),
        total: total - 1,
      }));
      return nextTick();
    });

    try {
      await http.delete(['/schemas', deletingSchema.id]);
      list.commitOptimisticUpdate();
    } catch (error) {
      routeTransition.start(() => {
        list.rollbackOptimisticUpdate();
        return nextTick();
      });
      throw error;
    }
  }

  async function copySchema(copyingSchema: ISchema): Promise<SchemaListItem> {
    return http.post(['/schemas', copyingSchema.id, 'copy'], {});
  }

  return {
    schemas,
    hasSchemas,
    load: list.load,
    createSchema,
    deleteSchema,
    copySchema,
  };
});
