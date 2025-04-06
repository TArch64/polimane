import { defineStore } from 'pinia';
import { useAsyncData, useHttpClient } from '@/composables';
import type { ISchema } from '@/models';

export const useSchemasStore = defineStore('schemas', () => {
  const httpClient = useHttpClient();

  const schemas = useAsyncData({
    loader: () => httpClient.get<ISchema[]>('/schemas'),
    default: [],
  });

  return { schemas };
});
