import type { IListFolder } from '../homeStore';

export interface IFolderAddSchemasInput {
  schemaIds: string[];
  folderId: string | null;
  folderName: string | null;
}

export interface IFolderAddSchemaStrategy {
  getFolders: () => IListFolder[];

  do(input: IFolderAddSchemasInput): Promise<void>;
}
