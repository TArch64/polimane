import { defineStore } from 'pinia';
import { toRef } from 'vue';
import { type HttpBody, useAsyncData, useHttpClient } from '@/composables';
import type { IFolder } from '@/models';
import type { IFolderAddSchemasInput } from './strategies';
import type { IListFolder } from './homeStore';

interface IFolderAddSchemasRequest {
  schemaIds: string[];
}

interface IFolderCreateRequest extends IFolderAddSchemasRequest {
  name: string;
  asList: boolean;
}

export interface IFolderAddSchemasOptions {
  asList?: boolean;
}

export interface IFolderAddSchemasWithScreenshotOptions extends IFolderAddSchemasOptions {
  asList: true;
}

export interface IFolderAddSchemasOutput<F extends IFolder> {
  newFolder?: F;
}

export const useHomeFoldersStore = defineStore('home/folders', () => {
  const http = useHttpClient();

  const list = useAsyncData({
    async loader() {
      return http.get<IFolder[]>('/folders');
    },

    once: true,
    default: [],
  });

  async function addSchemasToNewFolder(request: IFolderCreateRequest): Promise<IFolderAddSchemasOutput<IFolder | IListFolder>> {
    const folder = await http.post<IListFolder, IFolderCreateRequest>('/folders', request);

    list.data = [
      {
        id: folder.id,
        name: folder.name,
      },
      ...list.data,
    ];

    return { newFolder: folder };
  }

  async function addSchemasToExistingFolder(folderId: string, request: IFolderAddSchemasRequest): Promise<IFolderAddSchemasOutput<IFolder | IListFolder>> {
    await http.post<HttpBody, IFolderAddSchemasRequest>(['/folders', folderId, 'schemas'], request);
    return {};
  }

  async function addSchemas(input: IFolderAddSchemasInput, options: IFolderAddSchemasWithScreenshotOptions): Promise<IFolderAddSchemasOutput<IListFolder>>;
  async function addSchemas(input: IFolderAddSchemasInput, options?: IFolderAddSchemasOptions): Promise<IFolderAddSchemasOutput<IFolder>>;
  async function addSchemas(input: IFolderAddSchemasInput, options: IFolderAddSchemasOptions = {}): Promise<IFolderAddSchemasOutput<IFolder | IListFolder>> {
    const addRequest: IFolderAddSchemasRequest = {
      schemaIds: input.schemaIds,
    };

    if (input.folderId) {
      return addSchemasToExistingFolder(input.folderId, addRequest);
    }

    return addSchemasToNewFolder({
      ...addRequest,
      name: input.folderName!,
      asList: options.asList ?? false,
    });
  }

  return {
    load: list.load,
    invalidate: list.reset,
    folders: toRef(list, 'data'),
    addSchemas,
  };
});
