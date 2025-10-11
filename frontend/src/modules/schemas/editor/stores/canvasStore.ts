import { defineStore } from 'pinia';
import { ref } from 'vue';
import { clamp } from '@vueuse/core';

export const MIN_SCALE = 0.5;
export const MAX_SCALE = 10;

export const useCanvasStore = defineStore('schemas/editor/canvas', () => {
  const scale = ref(1);

  function setScale(value: number) {
    scale.value = clamp(value, MIN_SCALE, MAX_SCALE);
  }

  return {
    scale,
    setScale,
  };
});
