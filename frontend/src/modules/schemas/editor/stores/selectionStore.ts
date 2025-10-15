import { defineStore } from 'pinia';
import { ref } from 'vue';
import type { SchemaBeadCoord } from '@/models';
import { useSelectionArea, useSelectionResize } from './composables';

export interface IBeadSelection {
  from: SchemaBeadCoord;
  to: SchemaBeadCoord;
}

export const useSelectionStore = defineStore('schemas/editor/selection', () => {
  const isSelecting = ref(false);
  const toggleSelecting = (value: boolean) => isSelecting.value = value;

  const selected = ref<IBeadSelection | null>(null);
  const setSelected = (value: IBeadSelection | null) => selected.value = value;

  const area = useSelectionArea();
  const resize = useSelectionResize({ area, selected });

  return {
    area,
    resize,
    isSelecting,
    toggleSelecting,
    selected,
    setSelected,
  };
});
