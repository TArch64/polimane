import { defineStore } from 'pinia';
import { computed } from 'vue';
import { type IFolderAddSchemasInput, useHomeFoldersStore } from '@/modules/home/stores';
import { useHomeListStore } from './homeListStore';
import { useSchemasStore } from './schemasStore';

export const useFoldersStore = defineStore('home/list/folders', () => {
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

  return {
    folders,
    hasFolders,
    addSchemas,
  };
});
