import { defineStore } from 'pinia';
import { computed } from 'vue';
import { type HttpBody, useHttpClient } from '@/composables';
import { useHomeListStore } from './homeListStore';

export interface IFolderAddSchemasInput {
  schemaIds: string[];
  folderId: string | null;
  folderName: string | null;
}

interface IFolderCreateRequest {
  name: string;
  schemaIds: string[];
}

interface IFolderAddSchemasRequest {
  schemaIds: string[];
}

export const useFoldersStore = defineStore('home/list/folders', () => {
  const http = useHttpClient();
  const listStore = useHomeListStore();

  const folders = computed(() => listStore.list.data.folders);
  const hasFolders = computed(() => !!folders.value.length);

  async function addSchemasToNewFolder(name: string, schemaIds: string[]): Promise<void> {
    await http.post<HttpBody, IFolderCreateRequest>('/folders', {
      name,
      schemaIds,
    });
  }

  async function addSchemasToExistingFolder(folderId: string, schemaIds: string[]): Promise<void> {
    await http.post<HttpBody, IFolderAddSchemasRequest>(['/folders', folderId, 'schemas'], {
      schemaIds,
    });
  }

  async function addSchemas(input: IFolderAddSchemasInput): Promise<void> {
    input.folderId
      ? await addSchemasToExistingFolder(input.folderId, input.schemaIds)
      : await addSchemasToNewFolder(input.folderName!, input.schemaIds);

    const schemaIdsSet = new Set(input.schemaIds);
    listStore.list.data.total += 1 - input.schemaIds.length;
    listStore.list.data.schemas = listStore.list.data.schemas.filter((schema) => !schemaIdsSet.has(schema.id));
  }

  return {
    folders,
    hasFolders,
    addSchemas,
  };
});
