import { computed, ref, type Ref } from 'vue';
import { createAnimatedFrame } from '@/helpers';
import {
  type BeadCoord,
  getBeadSettings,
  type IPoint,
  isRefBead,
  isSpannableBead,
  Point,
  type SchemaBead,
  type SchemaSpannableBead,
  serializeBeadPoint,
} from '@/models';
import { Direction, isHorizontalDirection, isVerticalDirection } from '@/enums';
import { PaintEffect, useBeadsStore, useEditorStore, useToolsStore } from '../../stores';
import type { IBeadToolsOptions } from './IBeadToolsOptions';
import { useBeadCoord } from './useBeadCoord';
import { useBeadFactory } from './useBeadFactory';

export interface IBeadPaintingListeners {
  mousedown: (event: MouseEvent) => void;
  mousemove?: (event: MouseEvent) => void;
}

interface ISpanningBead {
  coord: BeadCoord;
  point: Point;
  original: SchemaSpannableBead;
  direction?: Direction;
}

export function useBeadPainting(options: IBeadToolsOptions): Ref<IBeadPaintingListeners> {
  const toolsStore = useToolsStore();
  const beadsStore = useBeadsStore();
  const editorStore = useEditorStore();

  const beadCoord = useBeadCoord(options);
  const beadFactory = useBeadFactory();
  const isPainting = ref(false);
  let spanning: ISpanningBead | null = null;

  function restrictSpanningPoint(refPoint: IPoint): void {
    const { direction, point } = spanning!;
    refPoint.x = isHorizontalDirection(direction!) ? refPoint.x : point.x;
    refPoint.y = isVerticalDirection(direction!) ? refPoint.y : point.y;
  }

  function detectDirection(from: IPoint, to: IPoint): Direction {
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

  function updateSpanningSize(point: IPoint) {
    getBeadSettings(spanning!.original).span = {
      x: point.x - spanning!.point.x,
      y: point.y - spanning!.point.y,
    };
  }

  const paint = createAnimatedFrame((event: MouseEvent, color: string | null) => {
    const point = beadCoord.getFromEvent(event);
    if (!point) return;

    let coord = serializeBeadPoint(point);
    let bead: SchemaBead | null;

    if (spanning) {
      if (spanning.point.isEqual(point)) {
        return;
      }

      if (spanning.direction) {
        restrictSpanningPoint(point);

        if (detectDirection(spanning.point, point) !== spanning.direction) {
          return;
        }

        coord = serializeBeadPoint(point);
      } else {
        spanning.direction = detectDirection(spanning.point, point);
      }

      if (isSameSpanningRef(coord)) {
        return;
      }

      const spanningCoord = serializeBeadPoint(spanning.point);
      bead = beadFactory.createRef(spanningCoord);
      updateSpanningSize(point);
    } else {
      const kind = toolsStore.activeBead;
      bead = beadFactory.create(kind, color);

      if (!!bead && isSpannableBead(bead)) {
        spanning = {
          coord,
          point: new Point(point),
          original: bead,
        };
      }
    }

    if (beadsStore.paint(coord, bead) === PaintEffect.EXTENDED) {
      beadCoord.clearCache();
    }
  });

  function onMouseup() {
    isPainting.value = false;
    spanning = null;
    beadCoord.clearCache();
  }

  function onMousedown(event: MouseEvent) {
    if (event.buttons === 1) {
      isPainting.value = true;
      paint(event, toolsStore.isEraser ? null : toolsStore.activeColor);
      addEventListener('mouseup', onMouseup, { once: true });
    }

    if (event.buttons === 2) {
      paint(event, null);
    }
  }

  function onMousemove(event: MouseEvent) {
    if (event.shiftKey || !isPainting.value) return;
    paint(event, toolsStore.isEraser ? null : toolsStore.activeColor);
  }

  return computed(() => ({
    mousedown: onMousedown,
    ...(isPainting.value ? { mousemove: onMousemove } : {}),
  }));
}
