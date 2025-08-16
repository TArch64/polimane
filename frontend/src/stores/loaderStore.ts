import { computed, ref } from 'vue';
import { defineStore } from 'pinia';

export interface ILoader {
  show: () => void;
  hide: () => void;
}

export const useLoaderStore = defineStore('loader', () => {
  const counter = ref(0);
  const isDisplaying = computed(() => counter.value > 0);

  const show = () => counter.value++;
  const hide = () => counter.value = Math.max(0, counter.value - 1);

  return {
    isDisplaying,
    show,
    hide,
  };
});
