export interface IFolderAddSchemasInput {
  schemaIds: string[];
  folderId: string | null;
  folderName: string | null;
}

export interface IFolderAddSchemaAdapter {
  do(input: IFolderAddSchemasInput): Promise<void>;
}
