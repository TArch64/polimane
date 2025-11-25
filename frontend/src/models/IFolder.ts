export interface IFolder {
  id: string;
  name: string;
}

export type FolderUpdate = Partial<Omit<IFolder, 'id'>>;
