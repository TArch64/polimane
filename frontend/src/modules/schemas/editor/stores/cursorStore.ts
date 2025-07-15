import Konva from 'konva';
import { computed, readonly, ref } from 'vue';
import { defineStore } from 'pinia';

export const useCursorStore = defineStore('schemas/editor/cursor', () => {
  const isPainting = ref(false);
  const isDragging = ref(false);

  function handleKeyDown(event: KeyboardEvent): boolean {
    if (event.metaKey) {
      isDragging.value = true;
      return true;
    }

    return false;
  }

  function handleKeyUp(event: KeyboardEvent) {
    if (!event.metaKey) {
      isDragging.value = false;
    }
  }

  function handleMouseDown(event: Konva.KonvaEventObject<MouseEvent>) {
    if (event.evt.buttons > 1) return;
    isPainting.value = event.evt.buttons === 1;
  }

  function handleMouseUp(event: Konva.KonvaEventObject<MouseEvent>) {
    isPainting.value = false;
  }

  return {
    handleKeyDown,
    handleKeyUp,
    handleMouseDown,
    handleMouseUp,
    isPainting: computed(() => isPainting.value && !isDragging.value),
    isDragging: readonly(isDragging),
  };
});
