import { computed, ref } from 'vue';
import { defineStore } from 'pinia';
import { useAsyncData, useHttpClient } from '@/composables';
import type { ListSchema } from '@/modules/home/stores';

const PAGINATION_PAGE = 100;

interface IListResponse {
  schemas: ListSchema[];
  total: number;
}

type ListRequestParams = {
  folder: string;
  offset: number;
  limit: number;
};

export const useFolderSchemasStore = defineStore('home/folder/schemas', () => {
  const http = useHttpClient();
  const folderId = ref<string>('');

  const list = useAsyncData({
    async loader(current): Promise<IListResponse> {
      return http.get<IListResponse, ListRequestParams>(['/schemas'], {
        folder: folderId.value,
        limit: PAGINATION_PAGE,
        offset: current.schemas.length,
      });
    },

    default: {
      schemas: [],
      total: 0,
    },
  });

  const schemas = computed(() => list.data.schemas);
  const canLoadNext = computed(() => schemas.value.length < list.data.total);

  async function load(id: string): Promise<void> {
    folderId.value = id;
    await list.load();
  }

  async function loadNext(): Promise<void> {
    if (list.isLoading) return;
    if (!canLoadNext.value) return;
    return list.load();
  }

  return {
    schemas,
    canLoadNext,
    load,
    loadNext,
  };
});
