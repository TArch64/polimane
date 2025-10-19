import { defineStore } from 'pinia';
import {
  type BeadCoord,
  getBeadSettings,
  type IPoint,
  parseBeadCoord,
  type SchemaBead,
  type SchemaBeads,
  type SchemaSpannableBead,
  serializeBeadCoord,
} from '@/models';
import { BeadKind, Direction, isBeadSpannableKind } from '@/enums';
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

  function isInArea(coord: IPoint, from: IPoint, to: IPoint): boolean {
    return coord.x >= from.x && coord.x <= to.x
      && coord.y >= from.y && coord.y <= to.y;
  }

  function getInArea(from: IPoint, to: IPoint): SchemaBeads {
    const entries = getObjectEntries(editorStore.schema.beads).filter(([coord]) => {
      return isInArea(parseBeadCoord(coord), from, to);
    });

    return Object.fromEntries(entries);
  }

  function onBeadBeforeRemove(coord: BeadCoord): void {
    const bead = editorStore.schema.beads[coord];
    if (!bead || !isBeadSpannableKind(bead.kind)) return;

    const point = parseBeadCoord(coord);
    const settings = getBeadSettings(bead as SchemaSpannableBead);

    const spanX = point.x + settings.span.x;
    const spanY = point.y + settings.span.y;

    const from: IPoint = {
      x: Math.min(point.x, spanX),
      y: Math.min(point.y, spanY),
    };

    const to: IPoint = {
      x: Math.max(point.x, spanX),
      y: Math.max(point.y, spanY),
    };

    for (let x = from.x; x <= to.x; x++) {
      for (let y = from.y; y <= to.y; y++) {
        const checkingCoord = serializeBeadCoord(x, y);
        const checkingBead = editorStore.schema.beads[checkingCoord];

        if (checkingBead?.kind === BeadKind.REF) {
          delete editorStore.schema.beads[checkingCoord];
        }
      }
    }
  }

  function remove(coord: BeadCoord): void {
    onBeadBeforeRemove(coord);
    delete editorStore.schema.beads[coord];
  }

  function paint(coord: BeadCoord, color: SchemaBead | null): PaintEffect | null {
    const currentColor = getColor(coord);

    if (currentColor === color) {
      return null;
    }

    if (color) {
      if (coord in editorStore.schema.beads) {
        onBeadBeforeRemove(coord);
      }

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

  return { paint, remove, getInArea };
});
