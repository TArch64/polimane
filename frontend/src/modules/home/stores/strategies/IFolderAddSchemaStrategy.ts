export interface IFolderAddSchemasInput {
  schemaIds: string[];
  folderId: string | null;
  folderName: string | null;
}

export interface IFolderAddSchemaStrategy {
  do(input: IFolderAddSchemasInput): Promise<void>;
}
