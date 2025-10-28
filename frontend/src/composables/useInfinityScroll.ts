import { nextTick, type Ref } from 'vue';
import { useWindowScroll } from '@vueuse/core';

export interface IInfinityScrollOptions {
  load: () => Promise<void>;
  canLoadNext: Ref<boolean>;
}

export function useInfinityScroll(options: IInfinityScrollOptions) {
  const { measure } = useWindowScroll({
    throttle: 100,
    observe: { mutation: false },

    onScroll: async () => {
      if (!options.canLoadNext.value) return;
      const scrollHeight = document.documentElement.scrollHeight;
      const distanceToEnd = scrollHeight - (window.scrollY + window.innerHeight);
      const threshold = Math.max(window.innerHeight * 0.75, 300);

      if (distanceToEnd <= threshold) {
        await options.load();
        await nextTick();
        measure();
      }
    },
  });
}
