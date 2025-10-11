import { defineStore } from 'pinia';
import { computed, reactive } from 'vue';
import type { INodeRect } from '@/models';
import { useCanvasStore } from './canvasStore';

export const useSelectionStore = defineStore('schemas/editor/selection', () => {
  const canvasStore = useCanvasStore();

  const selection = reactive<INodeRect>({
    x: 0,
    y: 0,
    width: 0,
    height: 0,
  });

  function setPoint(x: number, y: number) {
    selection.x = x;
    selection.y = y;
    selection.width = 0;
    selection.height = 0;
  }

  function reset() {
    selection.x = 0;
    selection.y = 0;
    selection.width = 0;
    selection.height = 0;
  }

  function extend(x: number, y: number) {
    selection.width += x;
    selection.height += y;
  }

  const resolvedSelection = computed(() => ({
    x: selection.width < 0 ? selection.x + selection.width : selection.x,
    y: selection.height < 0 ? selection.y + selection.height : selection.y,
    width: Math.abs(selection.width) * canvasStore.scale,
    height: Math.abs(selection.height) * canvasStore.scale,
  }));

  const isEmpty = computed(() => {
    const values = Object.values(resolvedSelection.value);
    return values.some((v) => !v);
  });

  return {
    selection: resolvedSelection,
    isEmpty,
    setPoint,
    reset,
    extend,
  };
});
