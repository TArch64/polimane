import type { FolderUpdate, IFolder } from '@/models';

export interface IFolderUpdateAdapter {
  do(folder: IFolder, update: FolderUpdate): Promise<void>;
}
