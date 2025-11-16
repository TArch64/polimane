import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { useAsyncData, useHttpClient } from '@/composables';
import type { IFolder } from '@/models';
import type { ListSchema } from '@/modules/home/stores';

export interface IFolderDetails extends IFolder {
  schemas: ListSchema[];
}

export const useFolderStore = defineStore('home/folder', () => {
  const http = useHttpClient();
  const folderId = ref<string | null>(null);

  const details = useAsyncData({
    async loader() {
      if (!folderId.value) return null;
      return http.get<IFolderDetails>(['/folders', folderId.value]);
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
