import { defineStore } from 'pinia';
import { computed } from 'vue';
import { useAsyncData, useHttpClient } from '@/composables';
import type { IListFolder, ListSchema } from '@/modules/home/stores';

const PAGINATION_PAGE = 100;

interface IListResponse {
  folders: IListFolder[];
  schemas: ListSchema[];
  total: number;
}

type ListRequestParams = {
  offset: number;
  limit: number;
};

export const useHomeListStore = defineStore('home/list', () => {
  const http = useHttpClient();

  const list = useAsyncData({
    async loader(current): Promise<IListResponse> {
      const response = await http.get<IListResponse, ListRequestParams>('/schemas', {
        limit: PAGINATION_PAGE,
        offset: current.folders.length + current.schemas.length,
      });

      return {
        folders: [...current.folders, ...response.folders],
        schemas: [...current.schemas, ...response.schemas],
        total: response.total,
      };
    },

    default: {
      folders: [],
      schemas: [],
      total: 0,
    },
  });

  const loadedTotal = computed(() => list.data.folders.length + list.data.schemas.length);
  const canLoadNext = computed(() => loadedTotal.value < list.data.total);

  const hasData = computed(() => {
    return !!loadedTotal.value || list.isLoading;
  });

  function load(): Promise<void> {
    return list.load(true);
  }

  async function loadNext(): Promise<void> {
    if (list.isLoading) return;
    if (!canLoadNext.value) return;
    return list.load();
  }

  return {
    list,
    hasData,
    canLoadNext,
    load,
    loadNext,
  };
});
