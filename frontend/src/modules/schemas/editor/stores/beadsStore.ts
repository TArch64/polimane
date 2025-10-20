import { defineStore } from 'pinia';
import {
  type BeadCoord,
  getBeadSettings,
  type IPoint,
  isRefBead,
  isSpannableBead,
  parseBeadCoord,
  type SchemaBead,
  type SchemaBeads,
  serializeBeadCoord,
} from '@/models';
import { BeadKind, Direction } from '@/enums';
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

  function* iterateArea(from: IPoint, to: IPoint): Generator<[BeadCoord, SchemaBead]> {
    for (let x = from.x; x <= to.x; x++) {
      for (let y = from.y; y <= to.y; y++) {
        const coord = serializeBeadCoord(x, y);
        const bead = editorStore.schema.beads[coord];
        if (bead) yield [coord, bead];
      }
    }
  }

  function getInArea(from: IPoint, to: IPoint): SchemaBeads {
    const beads: SchemaBeads = {};

    for (const [coord, bead] of iterateArea(from, to)) {
      beads[coord] = bead;
    }

    return beads;
  }

  function onBeadBeforeRemove(coord: BeadCoord): void {
    const bead = editorStore.schema.beads[coord];
    if (!bead) return;

    if (isRefBead(bead)) {
      return remove(getBeadSettings(bead).to);
    }

    if (isSpannableBead(bead)) {
      const { x, y } = parseBeadCoord(coord);
      const { span } = getBeadSettings(bead);

      const spanX = x + span.x;
      const spanY = y + span.y;

      const from: IPoint = {
        x: Math.min(x, spanX),
        y: Math.min(y, spanY),
      };

      const to: IPoint = {
        x: Math.max(x, spanX),
        y: Math.max(y, spanY),
      };

      for (const [coord, bead] of iterateArea(from, to)) {
        if (bead.kind === BeadKind.REF) {
          delete editorStore.schema.beads[coord];
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
