import { defineStore } from 'pinia';
import { computed } from 'vue';
import { type HttpBody, useHttpClient } from '@/composables';
import type { IFolderAddSchemasInput, IListFolder } from '@/modules/home/stores';
import { useHomeListStore } from './homeListStore';
import { useSchemasStore } from './schemasStore';

interface IFolderAddSchemasRequest {
  schemaIds: string[];
}

interface IFolderCreateRequest extends IFolderAddSchemasRequest {
  name: string;
}

export const useFoldersStore = defineStore('home/list/folders', () => {
  const http = useHttpClient();
  const listStore = useHomeListStore();
  const schemasStore = useSchemasStore();

  const folders = computed(() => listStore.list.data.folders);
  const hasFolders = computed(() => !!folders.value.length);

  async function addSchemasToNewFolder(name: string, addRequest: IFolderAddSchemasRequest): Promise<void> {
    const folder = await http.post<IListFolder, IFolderCreateRequest>('/folders', {
      name,
      ...addRequest,
    });

    listStore.list.data.total++;
    listStore.list.data.folders = [folder, ...listStore.list.data.folders];
  }

  async function addSchemasToExistingFolder(folderId: string, request: IFolderAddSchemasRequest): Promise<void> {
    await http.post<HttpBody, IFolderAddSchemasRequest>(['/folders', folderId, 'schemas'], request);
  }

  async function addSchemas(input: IFolderAddSchemasInput): Promise<void> {
    const addRequest: IFolderAddSchemasRequest = {
      schemaIds: input.schemaIds,
    };

    input.folderId
      ? await addSchemasToExistingFolder(input.folderId, addRequest)
      : await addSchemasToNewFolder(input.folderName!, addRequest);

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
