import type { IFolder } from '@/models';

export interface IFolderAddSchemasInput {
  schemaIds: string[];
  folderId: string | null;
  folderName: string | null;
}

export interface IFolderAddSchemaStrategy {
  getFolders: () => IFolder[];

  do(input: IFolderAddSchemasInput): Promise<void>;
}
