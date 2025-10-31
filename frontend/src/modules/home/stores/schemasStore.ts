import { defineStore } from 'pinia';
import { computed, nextTick, ref, type Ref, toRef } from 'vue';
import { type HttpBody, useAsyncData, useHttpClient, useRouteTransition } from '@/composables';
import type { ISchema, SchemaUpdate } from '@/models';
import { AccessLevel } from '@/enums';

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

interface IDeleteManySchemasBody {
  ids: string[];
}

export const useSchemasStore = defineStore('schemas/list', () => {
  const routeTransition = useRouteTransition();
  const http = useHttpClient();

  const selected: Ref<Set<string>> = ref(new Set());
  const clearSelection = () => selected.value = new Set();

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

  async function deleteMany(ids: Set<string>): Promise<void> {
    const idToAccess = Object.fromEntries(
      schemas.value.map((schema) => [schema.id, schema.access]),
    );

    for (const id of Array.from(ids)) {
      if (idToAccess[id] !== AccessLevel.ADMIN) {
        ids.delete(id);
      }
    }

    routeTransition.start(() => {
      list.makeOptimisticUpdate(({ list, total }) => ({
        list: list.filter((schema) => !ids.has(schema.id)),
        total: total - ids.size,
      }));

      return nextTick();
    });

    try {
      await http.delete<HttpBody, IDeleteManySchemasBody>(['/schemas', 'delete-many'], {
        ids: Array.from(ids),
      });

      list.commitOptimisticUpdate();

      if (canLoadNext.value && schemas.value.length < PAGINATION_PAGE) {
        await loadNext();
      }
    } catch (error) {
      routeTransition.start(() => {
        list.rollbackOptimisticUpdate();
        return nextTick();
      });
      throw error;
    }
  }

  async function deleteSchema(deletingSchema: SchemaListItem): Promise<void> {
    return deleteMany(new Set([deletingSchema.id]));
  }

  async function copySchema(copyingSchema: SchemaListItem): Promise<SchemaListItem> {
    const item = await http.post<SchemaListItem, HttpBody>(['/schemas', copyingSchema.id, 'copy'], {});
    list.data.total++;
    return item;
  }

  async function updateSchema(updatingSchema: SchemaListItem, patch: SchemaUpdate): Promise<void> {
    await http.patch<HttpBody, SchemaUpdate>(['/schemas', updatingSchema.id], patch);
    Object.assign(updatingSchema, patch);
  }

  return {
    schemas,
    hasSchemas,
    isLoading: toRef(list, 'isLoading'),
    canLoadNext,
    load,
    loadNext,
    selected,
    clearSelection,
    createSchema,
    deleteSchema,
    deleteMany,
    copySchema,
    updateSchema,
  };
});
