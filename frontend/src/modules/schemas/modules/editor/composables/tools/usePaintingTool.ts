import { computed, ref } from 'vue';
import { createAnimatedFrame, type ObjectEntry } from '@/helpers';
import {
  type BeadCoord,
  getBeadSettings,
  isRefBead,
  isSpannableBead,
  Point,
  type SchemaBead,
  type SchemaBeads,
  type SchemaSpannableBead,
  serializeBeadPoint,
} from '@/models';
import { Direction, isHorizontalDirection, isVerticalDirection } from '@/enums';
import {
  PaintEffect,
  useBeadFactory,
  useBeadsStore,
  useEditorStore,
  useToolsStore,
} from '../../stores';
import type { IEditorTool, IEditorToolOptions } from './tool';
import { useBeadCoord } from './useBeadCoord';

interface ISpanningBead {
  coord: BeadCoord;
  point: Point;
  original: SchemaSpannableBead;
  direction?: Direction;
}

export const usePaintingTool = (options: IEditorToolOptions) => {
  const toolsStore = useToolsStore();
  const beadsStore = useBeadsStore();
  const editorStore = useEditorStore();

  const beadCoord = useBeadCoord(options);
  const beadFactory = useBeadFactory();
  const isPainting = ref(false);
  let spanning: ISpanningBead | null = null;
  let lastPaintedPoint: Point | null = null;

  function restrictSpanningPoint(refPoint: Point): void {
    const { direction, point } = spanning!;
    refPoint.x = isHorizontalDirection(direction!) ? refPoint.x : point.x;
    refPoint.y = isVerticalDirection(direction!) ? refPoint.y : point.y;
  }

  function detectDirection(from: Point, to: Point): Direction {
    const deltaX = to.x - from.x;
    const deltaY = to.y - from.y;

    if (Math.abs(deltaX) > Math.abs(deltaY)) {
      return deltaX > 0 ? Direction.RIGHT : Direction.LEFT;
    } else {
      return deltaY > 0 ? Direction.BOTTOM : Direction.TOP;
    }
  }

  function isSameSpanningRef(coord: BeadCoord): boolean {
    const existingBead = editorStore.schema.beads[coord];

    return !!existingBead
      && isRefBead(existingBead)
      && getBeadSettings(existingBead).to === spanning!.coord;
  }

  function updateSpanningSize(point: Point) {
    getBeadSettings(spanning!.original).span = {
      x: point.x - spanning!.point.x,
      y: point.y - spanning!.point.y,
    };
  }

  function paintCell(point: Point, color: string): ObjectEntry<SchemaBeads> | null {
    let coord = serializeBeadPoint(point);
    let bead: SchemaBead | null;

    if (spanning) {
      if (spanning.point.isEqual(point)) {
        return null;
      }

      if (spanning.direction) {
        if (detectDirection(spanning.point, point) !== spanning.direction) {
          return null;
        }

        coord = serializeBeadPoint(point);
      } else {
        spanning.direction = detectDirection(spanning.point, point);
      }

      if (isSameSpanningRef(coord)) {
        return null;
      }

      const spanningCoord = serializeBeadPoint(spanning.point);
      bead = beadFactory.createRef(spanningCoord);
      updateSpanningSize(point);
    } else {
      const kind = toolsStore.activeBead;
      bead = beadFactory.create(kind, color);

      if (isSpannableBead(bead)) {
        spanning = {
          coord,
          point,
          original: bead,
        };
      }
    }

    return [coord, bead];
  }

  function buildSequence(current: Point): Point[] {
    if (!lastPaintedPoint) {
      return [current];
    }
    if (lastPaintedPoint.isEqual(current)) {
      return [];
    }
    const difference = lastPaintedPoint.getAxisDifference(current);
    if (difference.x > 0 && difference.y > 0) {
      return [current];
    }
    const axis = difference.x !== 0 ? 'x' : 'y';
    const baseX = Math.min(lastPaintedPoint.x, current.x);
    const baseY = Math.min(lastPaintedPoint.y, current.y);

    return Array.from({ length: difference[axis] + 1 }, (_, index): Point => (
      new Point(
        axis === 'x' ? baseX + index : current.x,
        axis === 'y' ? baseY + index : current.y,
      )
    ));
  }

  const paint = createAnimatedFrame((event: MouseEvent, color: string | null) => {
    const current = beadCoord.getFromEvent(event);
    if (!current) return;

    if (spanning?.direction) {
      restrictSpanningPoint(current);
    }

    const sequence = buildSequence(current);
    if (!sequence.length) return;

    if (!color) {
      const removing = sequence.map(serializeBeadPoint);
      beadsStore.removeMany(removing);
      lastPaintedPoint = current;
      return;
    }

    const paintEntries = sequence
      .map((point) => paintCell(point, color))
      .filter((entry) => entry !== null);

    if (!paintEntries.length) {
      return;
    }

    const effects = beadsStore.paintMany(Object.fromEntries(paintEntries));
    lastPaintedPoint = current;

    if (effects.has(PaintEffect.EXTENDED)) {
      beadCoord.clearCache();
    }
  });

  function onMouseup() {
    isPainting.value = false;
    spanning = null;
    lastPaintedPoint = null;
    beadCoord.clearCache();
  }

  function onMousedown(event: MouseEvent) {
    if (event.buttons === 0 || event.buttons > 2) {
      return;
    }

    switch (event.buttons) {
      case 1:
        isPainting.value = true;
        paint(event, toolsStore.isEraser ? null : toolsStore.activeColor);
        break;
      case 2:
        paint(event, null);
    }

    addEventListener('mouseup', onMouseup, { once: true });
  }

  function onMousemove(event: MouseEvent) {
    if (event.shiftKey || !isPainting.value) return;
    paint(event, toolsStore.isEraser ? null : toolsStore.activeColor);
  }

  return computed((): IEditorTool => ({
    level: 'content',

    listeners: {
      mousedown: onMousedown,
      ...(isPainting.value ? { mousemove: onMousemove } : {}),
    },
  }));
};
