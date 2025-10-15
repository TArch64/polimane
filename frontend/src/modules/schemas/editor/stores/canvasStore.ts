import { defineStore } from 'pinia';
import { ref } from 'vue';
import { clamp } from '@vueuse/core';
import type { IPoint } from '@/models';

export const MIN_SCALE = 0.5;
export const MAX_SCALE = 10;

export const useCanvasStore = defineStore('schemas/editor/canvas', () => {
  const scale = ref(1);
  const translation = ref<IPoint>({ x: 0, y: 0 });

  function setScale(value: number): number {
    scale.value = clamp(value, MIN_SCALE, MAX_SCALE);
    return scale.value;
  }

  function setTranslation(value: IPoint) {
    translation.value = value;
  }

  return {
    scale,
    setScale,
    translation,
    setTranslation,
  };
});
