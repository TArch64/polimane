import type { FileStorageEntryData, FileStorageEntryRaw } from './fileStorageSchema';

export class FileStorageEntry {
  static fromJSON(data: FileStorageEntryData): FileStorageEntry {
    return new FileStorageEntry(data.filename, data.title, data.updatedAt);
  }

  constructor(
    public filename: string,
    public title: string,
    public updatedAt: Date,
  ) {
  }

  toJSON(): FileStorageEntryRaw {
    return {
      filename: this.filename,
      title: this.title,
      updatedAt: this.updatedAt.getTime(),
    };
  }
}
