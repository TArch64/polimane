import type { NavigationGuardWithThis } from 'vue-router';
import { useStorageStore } from '@/stores';

export const checkStorageDirMiddleware: NavigationGuardWithThis<undefined> = async (to) => {
  const storageStore = useStorageStore();
  await storageStore.loadState();

  if (to.name === 'welcome') {
    return storageStore.isDirectorySelected ? { name: 'home' } : undefined;
  }

  return storageStore.isDirectorySelected ? undefined : { name: 'welcome' };
};
