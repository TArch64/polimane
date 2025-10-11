import { defineStore } from 'pinia';
import { computed, reactive, ref } from 'vue';
import type { INodeRect, SchemaBeadCoord } from '@/models';

export interface IBeadSelection {
  from: SchemaBeadCoord;
  to: SchemaBeadCoord;
}

export const useSelectionStore = defineStore('schemas/editor/selection', () => {
  const isSelecting = ref(false);
  const toggleSelecting = (value: boolean) => isSelecting.value = value;

  const selected = ref<IBeadSelection | null>(null);
  const setSelected = (value: IBeadSelection | null) => selected.value = value;

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
    width: Math.abs(selection.width),
    height: Math.abs(selection.height),
  }));

  return {
    selection: resolvedSelection,
    setPoint,
    reset,
    extend,
    isSelecting,
    toggleSelecting,
    selected,
    setSelected,
  };
});
