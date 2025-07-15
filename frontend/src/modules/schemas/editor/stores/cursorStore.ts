import Konva from 'konva';
import { readonly, ref } from 'vue';
import { defineStore } from 'pinia';

export const useCursorStore = defineStore('schemas/editor/cursor', () => {
  const isPainting = ref(false);

  function handleMouseDown(event: Konva.KonvaEventObject<MouseEvent>) {
    if (event.evt.buttons > 1) return;
    isPainting.value = event.evt.buttons === 1;
  }

  function handleMouseUp(event: Konva.KonvaEventObject<MouseEvent>) {
    isPainting.value = false;
  }

  return {
    handleMouseDown,
    handleMouseUp,
    isPainting: readonly(isPainting),
  };
});
