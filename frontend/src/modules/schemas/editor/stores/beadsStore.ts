import { defineStore } from 'pinia';
import { parseSchemaBeadCoord, type SchemaBeadCoord, type SchemaSizeDirection } from '@/models';
import { useEditorStore } from './editorStore';
import { usePaletteStore } from './paletteStore';

export const useBeadsStore = defineStore('schemas/editor/beads', () => {
  const editorStore = useEditorStore();
  const paletteStore = usePaletteStore();

  function getColor(coord: SchemaBeadCoord) {
    return editorStore.schema.beads[coord] ?? null;
  }

  function checkExtendingPaint(coord: SchemaBeadCoord): SchemaSizeDirection[] {
    const [x, y] = parseSchemaBeadCoord(coord);
    const size = editorStore.schema.size;
    const directions: SchemaSizeDirection[] = [];

    if (x <= 0) {
      if (size.left + x < 3) {
        directions.push('left');
      }
    } else {
      if (size.right - x < 3) {
        directions.push('right');
      }
    }

    if (y <= 0) {
      if (size.top + y < 3) {
        directions.push('top');
      }
    } else {
      if (size.bottom - y < 3) {
        directions.push('bottom');
      }
    }

    return directions;
  }

  function extendSchemaSize(directions: SchemaSizeDirection[]): void {
    for (const direction of directions) {
      editorStore.schema.size[direction] += 10;
    }
  }

  function paint(coord: SchemaBeadCoord) {
    if (paletteStore.activeColor) {
      editorStore.schema.beads[coord] = paletteStore.activeColor;

      const extendingDirections = checkExtendingPaint(coord);
      if (extendingDirections.length) extendSchemaSize(extendingDirections);
      return;
    }

    if (editorStore.schema.beads[coord]) {
      delete editorStore.schema.beads[coord];
    }
  }

  return { getColor, paint };
});
