import { defineStore } from 'pinia';
import {
  type BeadCoord,
  getBeadSettings,
  isRefBead,
  isSpannableBead,
  parseBeadCoord,
  Point,
  type SchemaBead,
  type SchemaBeads,
  type SchemaSpannableBead,
  serializeBeadCoord,
} from '@/models';
import { BeadKind, Direction } from '@/enums';
import { deleteObjectKeys, getObjectEntries } from '@/helpers';
import { useSchemaBeadsCounter } from '@/composables/subscription';
import { useEditorStore } from './editorStore';

export enum PaintEffect {
  EXTENDED = 'extend',
}

export const useBeadsStore = defineStore('schemas/editor/beads', () => {
  const editorStore = useEditorStore();
  const counter = useSchemaBeadsCounter(() => editorStore.schema);

  function getBead(coord: BeadCoord) {
    return editorStore.schema.beads[coord];
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

  function* iterateArea(from: Point, to: Point): Generator<[BeadCoord, SchemaBead]> {
    for (let x = from.x; x <= to.x; x++) {
      for (let y = from.y; y <= to.y; y++) {
        const coord = serializeBeadCoord(x, y);
        const bead = getBead(coord);
        if (bead) yield [coord, bead];
      }
    }
  }

  function getSpanBeads<K extends BeadKind = BeadKind>(
    coord: BeadCoord,
    bead: SchemaSpannableBead,
    out: SchemaBeads<K> = {},
  ): SchemaBeads<K> {
    const { x, y } = parseBeadCoord(coord);
    const settings = getBeadSettings(bead);

    const spanX = x + settings.span.x;
    const spanY = y + settings.span.y;

    const from = new Point(Math.min(x, spanX), Math.min(y, spanY));
    const to = new Point(Math.max(x, spanX), Math.max(y, spanY));

    for (const [spanCoord, spanBead] of iterateArea(from, to)) {
      if (coord !== spanCoord) {
        out[spanCoord] = spanBead as SchemaBead<K>;
      }
    }

    return out;
  }

  function getInArea(from: Point, to: Point): SchemaBeads {
    const beads: SchemaBeads = {};

    for (let [coord, bead] of iterateArea(from, to)) {
      if (isRefBead(bead)) {
        coord = getBeadSettings(bead).to;
        bead = getBead(coord)!;
      }

      if (coord in beads) {
        continue;
      }

      beads[coord] = bead;

      if (isSpannableBead(bead)) {
        getSpanBeads(coord, bead, beads);
      }
    }

    return beads;
  }

  function unsetBeads(removingSet: Set<BeadCoord>): void {
    editorStore.schema.beads = deleteObjectKeys(editorStore.schema.beads, removingSet);
  }

  function onBeadBeforeRemove(coord: BeadCoord, removingSet: Set<BeadCoord>): void {
    const bead = getBead(coord);
    if (!bead) return;

    if (isRefBead(bead)) {
      removingSet.add(getBeadSettings(bead).to);
      return;
    }

    if (isSpannableBead(bead)) {
      const spanBeads = getSpanBeads(coord, bead);

      for (const [coord, bead] of getObjectEntries(spanBeads)) {
        if (bead.kind === BeadKind.REF) {
          removingSet.add(coord);
        }
      }
    }
  }

  function removeInternal(coords: BeadCoord[]): void {
    const removingSet = new Set(coords);

    for (const coord of coords) {
      onBeadBeforeRemove(coord, removingSet);
    }

    unsetBeads(removingSet);
  }

  function remove(coord: BeadCoord): void {
    if (getBead(coord)) {
      removeInternal([coord]);
      counter.current--;
    }
  }

  function removeMany(coords: BeadCoord[]) {
    const existing = coords.filter(getBead);
    removeInternal(existing);
    counter.current -= existing.length;
  }

  function tryExtendSchema(coord: BeadCoord): PaintEffect | null {
    const extendingDirections = checkExtendingPaint(coord);

    if (!extendingDirections.length) {
      return null;
    }

    extendSchemaSize(extendingDirections);
    return PaintEffect.EXTENDED;
  }

  function paintMany(beads: SchemaBeads): Set<PaintEffect> {
    let localCounter = counter.current;
    const effects = new Set<PaintEffect>();

    const allowedSet = getObjectEntries(beads).filter(([coord, bead]) => {
      if (getBead(coord)) return true;
      if (localCounter >= counter.max!) return false;
      if (!isRefBead(bead)) localCounter++;
      return true;
    });

    if (allowedSet.length === 0) {
      return effects;
    }

    editorStore.schema.beads = {
      ...editorStore.schema.beads,
      ...Object.fromEntries(allowedSet),
    };

    counter.current = localCounter;

    for (const [coord] of allowedSet) {
      const newEffect = tryExtendSchema(coord);

      if (newEffect) {
        effects.add(newEffect);
      }
    }

    return effects;
  }

  return {
    paintMany,
    remove,
    removeMany,
    getInArea,
    getSpanBeads,
  };
});
