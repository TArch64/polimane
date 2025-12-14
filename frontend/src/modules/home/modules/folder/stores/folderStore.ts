import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { useAsyncData, useHttpClient } from '@/composables';
import type { FolderUpdate, IFolder } from '@/models';

export interface IFolderDeleteOptions {
  deleteSchemas: boolean;
}

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

  async function update(update: FolderUpdate) {
    await details.optimisticUpdate()
      .begin((state) => Object.assign(state!, update))
      .commit(async () => {
        await http.patch<FolderUpdate>(['/folders', folderId.value], update);
      });
  }

  async function deleteFolder(options: IFolderDeleteOptions) {
    await http.delete(['/folders', folderId.value], options);
  }

  return {
    folder: computed(() => details.data!),
    load,
    update,
    delete: deleteFolder,
  };
});
