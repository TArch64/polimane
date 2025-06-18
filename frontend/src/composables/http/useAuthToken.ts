import { useLocalStorage } from '@vueuse/core';

export function useAuthToken() {
  return useLocalStorage<string | undefined>('pa', undefined, {
    writeDefaults: false,
    shallow: true,
  });
}
