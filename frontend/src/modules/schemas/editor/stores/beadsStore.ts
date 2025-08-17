import { defineStore } from 'pinia';
import type { SchemaBeedCoord } from '@/models';
import { useEditorStore } from './editorStore';
import { usePaletteStore } from './paletteStore';

export const useBeadsStore = defineStore('schemas/editor/beads', () => {
  const editorStore = useEditorStore();
  const paletteStore = usePaletteStore();

  function getColor(coord: SchemaBeedCoord) {
    return editorStore.schema.beads[coord] ?? null;
  }

  function paint(coord: SchemaBeedCoord) {
    if (paletteStore.activeColor) {
      editorStore.schema.beads[coord] = paletteStore.activeColor;
      return;
    }

    if (editorStore.schema.beads[coord]) {
      delete editorStore.schema.beads[coord];
    }
  }

  return { getColor, paint };
});
