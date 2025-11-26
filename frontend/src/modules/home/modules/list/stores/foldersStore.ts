import { defineStore } from 'pinia';
import { computed } from 'vue';
import { type IFolderAddSchemasInput, useHomeFoldersStore } from '@/modules/home/stores';
import type { FolderUpdate, IFolder } from '@/models';
import { useHttpClient } from '@/composables';
import { useHomeListStore } from './homeListStore';
import { useSchemasStore } from './schemasStore';

export const useFoldersStore = defineStore('home/list/folders', () => {
  const http = useHttpClient();
  const homeFoldersStore = useHomeFoldersStore();
  const listStore = useHomeListStore();
  const schemasStore = useSchemasStore();

  const folders = computed(() => listStore.list.data.folders);
  const hasFolders = computed(() => !!folders.value.length);

  async function addSchemas(input: IFolderAddSchemasInput): Promise<void> {
    const result = await homeFoldersStore.addSchemas(input, { asList: true });

    if (result.newFolder) {
      listStore.list.data.total++;
      listStore.list.data.folders = [result.newFolder, ...listStore.list.data.folders];
    }

    const schemaIdsSet = new Set(input.schemaIds);
    listStore.list.data.total -= input.schemaIds.length;
    listStore.list.data.schemas = listStore.list.data.schemas.filter((schema) => !schemaIdsSet.has(schema.id));
    schemasStore.clearSelection();
  }

  async function updateFolder(folder: IFolder, update: FolderUpdate) {
    await listStore.list.optimisticUpdate()
      .begin(({ folders }) => {
        const index = folders.findIndex((f) => f.id === folder.id);
        folders[index] = { ...folders[index]!, ...update };
      })
      .commit(async () => {
        await http.patch<FolderUpdate>(['/folders', folder.id], update);
      });
  }

  async function deleteFolder(deleting: IFolder): Promise<void> {
    await listStore.list.optimisticUpdate()
      .inTransition()
      .begin((state) => {
        state.folders = state.folders.filter((folder) => folder.id !== deleting.id);
        state.total--;
      })
      .commit(async () => {
        await http.delete(['/folders', deleting.id]);
      });
  }

  return {
    folders,
    hasFolders,
    addSchemas,
    updateFolder,
    deleteFolder,
  };
});
