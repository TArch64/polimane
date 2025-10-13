import { defineStore } from 'pinia';
import { ref } from 'vue';
import type { SchemaBeadCoord } from '@/models';
import { useSelectionArea } from './composables';

export interface IBeadSelection {
  from: SchemaBeadCoord;
  to: SchemaBeadCoord;
}

export const useSelectionStore = defineStore('schemas/editor/selection', () => {
  const area = useSelectionArea();

  const isSelecting = ref(false);
  const toggleSelecting = (value: boolean) => isSelecting.value = value;

  const selected = ref<IBeadSelection | null>(null);
  const setSelected = (value: IBeadSelection | null) => selected.value = value;

  return {
    area,
    isSelecting,
    toggleSelecting,
    selected,
    setSelected,
  };
});
