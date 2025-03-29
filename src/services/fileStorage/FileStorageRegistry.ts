import { FileStorageEntry } from './FileStorageEntry';

export class FileStorageRegistry {
  static readonly SCHEMA_VERSION = 1;

  static async create(directoryHandle: FileSystemDirectoryHandle): Promise<FileStorageRegistry> {
    const registry = new FileStorageRegistry(directoryHandle);
    await registry.initialize();
    return registry;
  }

  private fileHandle!: FileSystemFileHandle;
  entries: FileStorageEntry[] = [];

  constructor(
    private readonly directoryHandle: FileSystemDirectoryHandle,
  ) {
  }

  private async initialize(): Promise<void> {
    this.fileHandle = await this.directoryHandle.getFileHandle('_registry.json', { create: true });
    await this.sync();
  }

  private async sync(): Promise<void> {
    const file = await this.fileHandle.getFile();

    if (!file.size) {
      await this.writeEntries();
      return;
    }
  }

  private async writeEntries(): Promise<void> {
    const json = JSON.stringify({
      version: FileStorageRegistry.SCHEMA_VERSION,
      updatedAt: Date.now(),
      entries: this.entries.map((entry) => entry.toJSON()),
    }, null, import.meta.env.DEV ? 2 : undefined);

    const writable = await this.fileHandle.createWritable({ keepExistingData: false });
    await writable.write(json);
    await writable.close();
  }
}
