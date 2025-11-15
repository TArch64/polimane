import { defineStore } from 'pinia';
import { computed } from 'vue';
import type { IFolder, ISchema } from '@/models';
import { useAsyncData, useHttpClient } from '@/composables';

const PAGINATION_PAGE = 100;

export type ListSchema = Omit<ISchema, 'beads' | 'size'>;

interface IListResponse {
  folders: IFolder[];
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
    loader: async (current): Promise<IListResponse> => {
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

  function load(): Promise<void> {
    list.reset();
    return list.load();
  }

  async function loadNext(): Promise<void> {
    if (list.isLoading) return;
    if (!canLoadNext.value) return;
    return list.load();
  }

  return {
    list,
    canLoadNext,
    load,
    loadNext,
  };
});
