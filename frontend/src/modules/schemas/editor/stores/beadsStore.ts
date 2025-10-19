import { defineStore } from 'pinia';
import {
  type BeadCoord,
  type IPoint,
  parseBeadCoord,
  type SchemaBead,
  type SchemaBeads,
} from '@/models';
import { Direction } from '@/enums';
import { getObjectEntries } from '@/helpers';
import { useEditorStore } from './editorStore';

export enum PaintEffect {
  EXTENDED = 'extend',
}

export const useBeadsStore = defineStore('schemas/editor/beads', () => {
  const editorStore = useEditorStore();

  function getColor(coord: BeadCoord) {
    return editorStore.schema.beads[coord] ?? null;
  }

  function checkExtendingPaint(coord: BeadCoord): Direction[] {
    const { x, y } = parseBeadCoord(coord);
    const size = editorStore.schema.size;
    const directions: Direction[] = [];

    if (x <= 0) {
      if (size.left + x < 3) {
        directions.push(Direction.LEFT);
      }
    } else {
      if (size.right - x < 3) {
        directions.push(Direction.RIGHT);
      }
    }

    if (y <= 0) {
      if (size.top + y < 3) {
        directions.push(Direction.TOP);
      }
    } else {
      if (size.bottom - y < 3) {
        directions.push(Direction.BOTTOM);
      }
    }

    return directions;
  }

  function extendSchemaSize(directions: Direction[]): void {
    for (const direction of directions) {
      editorStore.schema.size[direction] += 10;
    }
  }

  function remove(coord: BeadCoord): void {
    delete editorStore.schema.beads[coord];
  }

  function paint(coord: BeadCoord, color: SchemaBead | null): PaintEffect | null {
    const currentColor = getColor(coord);

    if (currentColor === color) {
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

    if (currentColor) {
      remove(coord);
    }

    return null;
  }

  function getInArea(from: IPoint, to: IPoint): SchemaBeads {
    const entries = getObjectEntries<SchemaBeads>(editorStore.schema.beads).filter(([coord]) => {
      const { x, y } = parseBeadCoord(coord);
      return x >= from.x && x <= to.x && y >= from.y && y <= to.y;
    });

    return Object.fromEntries(entries);
  }

  return { paint, remove, getInArea };
});
