import { useMediaQuery } from '@vueuse/core';

export function useMobileScreen() {
  return useMediaQuery('(max-width: 768px)');
}
