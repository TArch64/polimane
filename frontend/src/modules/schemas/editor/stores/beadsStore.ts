import { defineStore } from 'pinia';
import { computed } from 'vue';
import { useContrast } from '@editor/composables';
import { parseSchemaBeadCoord, type SchemaBeadCoord, type SchemaSizeDirection } from '@/models';
import { useEditorStore } from './editorStore';

const EMPTY_LIGHT = 'rgba(0, 0, 0, 0.1)';
const EMPTY_DARK = 'rgba(255, 255, 255, 0.1)';

export enum PaintEffect {
  EXTENDED = 'extend',
}

export const useBeadsStore = defineStore('schemas/editor/beads', () => {
  const editorStore = useEditorStore();

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

  function paint(coord: SchemaBeadCoord, color: string | null): PaintEffect | null {
    if (getColor(coord) === color) {
      return null;
    }

    if (color) {
      editorStore.schema.beads[coord] = color;

      const extendingDirections = checkExtendingPaint(coord);

      if (extendingDirections.length) {
        extendSchemaSize(extendingDirections);
        return PaintEffect.EXTENDED;
      }

      return null;
    }

    if (editorStore.schema.beads[coord]) {
      delete editorStore.schema.beads[coord];
    }

    return null;
  }

  const emptyContrast = useContrast(() => editorStore.schema.backgroundColor, EMPTY_DARK);
  const emptyColor = computed(() => emptyContrast.value < 4.5 ? EMPTY_LIGHT : EMPTY_DARK);

  return { paint, emptyColor };
});
