import { defineStore } from 'pinia';
import { IdbStorage } from '../services';
import { computed, ref } from 'vue';

const directoryHandleKey = IdbStorage.key<FileSystemDirectoryHandle>('storage-directory');

export const useStorageStore = defineStore('storage', () => {
  const directoryHandle = ref<FileSystemDirectoryHandle | null>(null);
  const isDirectorySelected = computed(() => !!directoryHandle.value);

  async function loadState(): Promise<void> {
    directoryHandle.value ??= await IdbStorage.instance.getItem(directoryHandleKey);
  }

  async function selectDirectory() {
    const handle = await window.showDirectoryPicker();
  }

  return { loadState, isDirectorySelected };
});
