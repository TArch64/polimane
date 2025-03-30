import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { FileStorageRegistry, IdbStorage } from '@/services';

const directoryHandleKey = IdbStorage.key<FileSystemDirectoryHandle>('storage-directory');

export const useStorageStore = defineStore('storage', () => {
  const directoryHandle = ref<FileSystemDirectoryHandle | null>(null);
  const isDirectorySelected = computed(() => !!directoryHandle.value);
  let registry: FileStorageRegistry;

  async function loadState(): Promise<void> {
    directoryHandle.value ??= await IdbStorage.instance.getItem(directoryHandleKey);

    if (!directoryHandle.value) {
      return;
    }

    const permission = await directoryHandle.value.queryPermission({ mode: 'readwrite' });

    if (permission !== 'granted') {
      directoryHandle.value = null;
      return;
    }

    registry = await FileStorageRegistry.create(directoryHandle.value);
  }

  async function selectDirectory(): Promise<void> {
    const handle = await showDirectoryPicker({
      id: 'polimane-storage',
      mode: 'readwrite',
      startIn: 'documents',
    });

    directoryHandle.value = handle;
    await IdbStorage.instance.setItem(directoryHandleKey, handle);
    registry = await FileStorageRegistry.create(directoryHandle.value!);
  }

  return { loadState, isDirectorySelected, selectDirectory };
});
