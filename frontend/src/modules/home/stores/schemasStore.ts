import { defineStore } from 'pinia';
import { computed, nextTick, toRef } from 'vue';
import { type HttpBody, useAsyncData, useHttpClient, useRouteTransition } from '@/composables';
import type { ISchema } from '@/models';

const PAGINATION_PAGE = 100;

export type SchemaListItem = Omit<ISchema, 'beads' | 'size'>;

interface IListResponse {
  list: SchemaListItem[];
  total: number;
}

type ListRequestParams = {
  offset: number;
  limit: number;
};

export interface ICreateSchemaRequest {
  name: string;
}

type UpdateSchemaRequest = Partial<Omit<ISchema, 'id'>>;

export const useSchemasStore = defineStore('schemas/list', () => {
  const routeTransition = useRouteTransition();
  const http = useHttpClient();

  const list = useAsyncData({
    loader: async (current): Promise<IListResponse> => {
      const response = await http.get<IListResponse, ListRequestParams>('/schemas', {
        limit: PAGINATION_PAGE,
        offset: current.list.length,
      });

      return {
        list: [...current.list, ...response.list],
        total: response.total,
      };
    },
    default: { list: [], total: 0 },
  });

  const schemas = computed(() => list.data.list);
  const hasSchemas = computed(() => !!schemas.value.length);
  const canLoadNext = computed(() => schemas.value.length < list.data.total);

  function load(): Promise<void> {
    list.reset();
    return list.load();
  }

  async function loadNext(): Promise<void> {
    if (list.isLoading) return;
    if (!canLoadNext.value) return;
    return list.load();
  }

  async function createSchema(input: ICreateSchemaRequest): Promise<SchemaListItem> {
    const item = await http.post<SchemaListItem, ICreateSchemaRequest>('/schemas', input);
    list.data.total++;
    return item;
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
    const item = await http.post<SchemaListItem, HttpBody>(['/schemas', copyingSchema.id, 'copy'], {});
    list.data.total++;
    return item;
  }

  async function updateSchema(updatingSchema: ISchema, patch: UpdateSchemaRequest): Promise<void> {
    await http.patch<HttpBody, UpdateSchemaRequest>(['/schemas', updatingSchema.id], patch);
    Object.assign(updatingSchema, patch);
  }

  return {
    schemas,
    hasSchemas,
    isLoading: toRef(list, 'isLoading'),
    canLoadNext,
    load,
    loadNext,
    createSchema,
    deleteSchema,
    copySchema,
    updateSchema,
  };
});
