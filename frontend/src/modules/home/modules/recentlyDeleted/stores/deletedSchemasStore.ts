import { computed, type Ref, ref, toRef } from 'vue';
import { defineStore } from 'pinia';
import { type HttpBody, useAsyncData, useHttpClient } from '@/composables';
import type { IDeleteManySchemasRequest, ListSchema } from '@/modules/home/stores';

const PAGINATION_PAGE = 100;

interface IListResponse {
  schemas: ListSchema[];
  total: number;
}

type ListRequestParams = {
  offset: number;
  limit: number;
};

export const useDeletedSchemasStore = defineStore('home/recently-deleted/schemas', () => {
  const http = useHttpClient();

  const list = useAsyncData({
    async loader(current): Promise<IListResponse> {
      const res = await http.get<IListResponse, ListRequestParams>(['/schemas/deleted'], {
        limit: PAGINATION_PAGE,
        offset: current.schemas.length,
      });

      return {
        schemas: [...current.schemas, ...res.schemas],
        total: res.total || current.total,
      };
    },

    default: {
      schemas: [],
      total: 0,
    },
  });

  const schemas = computed(() => list.data.schemas);
  const canLoadNext = computed(() => schemas.value.length < list.data.total);

  const selected: Ref<Set<string>> = ref(new Set());
  const clearSelection = () => selected.value = new Set();

  async function load(): Promise<void> {
    await list.load();
  }

  async function loadNext(): Promise<void> {
    if (list.isLoading) return;
    if (!canLoadNext.value) return;
    return list.load();
  }

  async function deleteMany(ids: string[]): Promise<void> {
    const idsSet = new Set(ids);

    await list.optimisticUpdate()
      .inTransition()
      .begin((state) => {
        state.schemas = state.schemas.filter((schema) => !idsSet.has(schema.id));
        state.total -= ids.length;
      })
      .commit(async () => {
        await http.delete<HttpBody, IDeleteManySchemasRequest>(['/schemas', 'delete-many'], { ids });
      });

    if (canLoadNext.value && schemas.value.length < PAGINATION_PAGE) {
      await loadNext();
    }
  }

  async function deleteSchema(deleting: ListSchema): Promise<void> {
    return deleteMany([deleting.id]);
  }

  return {
    schemas,
    selected,
    canLoadNext,
    isLoading: toRef(list, 'isLoading'),
    load,
    loadNext,
    clearSelection,
    deleteMany,
    deleteSchema,
  };
});
