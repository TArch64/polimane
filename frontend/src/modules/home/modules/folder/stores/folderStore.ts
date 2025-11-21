import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { useAsyncData, useHttpClient } from '@/composables';
import type { IFolder } from '@/models';

export const useFolderStore = defineStore('home/folder', () => {
  const http = useHttpClient();
  const folderId = ref<string>('');

  const details = useAsyncData({
    async loader() {
      return http.get<IFolder>(['/folders', folderId.value]);
    },
    default: null,
  });

  async function load(id: string): Promise<void> {
    folderId.value = id;
    await details.load();
  }

  return {
    load,
    folder: computed(() => details.data!),
  };
});
