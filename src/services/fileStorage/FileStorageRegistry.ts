import { parseAsync } from 'valibot';
import { FileStorageEntry } from './FileStorageEntry';
import type { FileStorageData, FileStorageRaw } from './fileStorageSchema';
import { fileStorageSchema } from './fileStorageSchema';

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
      await this.write();
      return;
    }

    const data = await this.read(file);
    this.entries = data.entries.map((entry) => FileStorageEntry.fromJSON(entry));
  }

  private async write(): Promise<void> {
    const jsonSpace = import.meta.env.DEV ? 2 : undefined;

    const json = JSON.stringify({
      version: FileStorageRegistry.SCHEMA_VERSION,
      updatedAt: Date.now(),
      entries: this.entries.map((entry) => entry.toJSON()),
    } satisfies FileStorageRaw, null, jsonSpace);

    const writable = await this.fileHandle.createWritable({ keepExistingData: false });
    await writable.write(import.meta.env.DEV ? json : btoa(json));
    await writable.close();
  }

  private async read(file: File): Promise<FileStorageData> {
    const text = await this.readFileText(file);
    const json = JSON.parse(import.meta.env.DEV ? text : atob(text));

    return parseAsync(fileStorageSchema(), json, {
      abortEarly: true,
    });
  }

  private readFileText(file: File): Promise<string> {
    return new Promise((resolve, reject) => {
      const reader = new FileReader();
      reader.onload = () => reader.result ? resolve(reader.result as string) : reject(reader.error);
      reader.onerror = () => reject(reader.error);
      reader.readAsText(file);
    });
  }
}
