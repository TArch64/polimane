import { useLocalStorage } from '@vueuse/core';

export function useAccessToken() {
  return useLocalStorage<string | undefined>('p.access', undefined, {
    writeDefaults: false,
    shallow: true,
  });
}

export function useRefreshAccessToken() {
  return useLocalStorage<string | undefined>('p.refresh', undefined, {
    writeDefaults: false,
    shallow: true,
  });
}
