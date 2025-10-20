import { defineStore } from 'pinia';
import { ref } from 'vue';
import { type BeadCoord, parseBeadCoord } from '@/models';
import { getObjectKeys } from '@/helpers';
import { useBeadsStore } from './beadsStore';
import { useSelectionArea, useSelectionResize } from './composables';

export interface IBeadSelection {
  from: BeadCoord;
  to: BeadCoord;
}

export const useSelectionStore = defineStore('schemas/editor/selection', () => {
  const beadsStore = useBeadsStore();

  const isSelecting = ref(false);
  const toggleSelecting = (value: boolean) => isSelecting.value = value;

  const selected = ref<IBeadSelection | null>(null);
  const setSelected = (value: IBeadSelection | null) => selected.value = value;

  const area = useSelectionArea();
  const resize = useSelectionResize({ area, selected });

  function reset() {
    setSelected(null);
    area.reset();
    resize.reset();
  }

  function removeSelected() {
    if (!selected.value) return;

    const from = parseBeadCoord(selected.value.from);
    const to = parseBeadCoord(selected.value.to);
    const selectedBeads = beadsStore.getInArea(from, to);

    for (const coord of getObjectKeys(selectedBeads)) {
      beadsStore.remove(coord);
    }

    reset();
  }

  return {
    area,
    resize,
    isSelecting,
    toggleSelecting,
    selected,
    setSelected,
    removeSelected,
    reset,
  };
});
