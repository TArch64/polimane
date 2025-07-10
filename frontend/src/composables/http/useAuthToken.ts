import { useLocalStorage } from '@vueuse/core';

export function useAuthToken() {
  return useLocalStorage<string | undefined>('p.auth', undefined, {
    writeDefaults: false,
    shallow: true,
  });
}
