import { computed, ref, type Ref } from 'vue';
import { createAnimatedFrame } from '@/helpers';
import {
  type BeadCoord,
  getBeadSettings,
  type IPoint,
  isRefBead,
  isSpannableBead,
  type SchemaBead,
  type SchemaSpannableBead,
  serializeBeadPoint,
} from '@/models';
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
  point: IPoint;
  original: SchemaSpannableBead;
  direction?: 'x' | 'y';
}

export function useBeadPainting(options: IBeadToolsOptions): Ref<IBeadPaintingListeners> {
  const toolsStore = useToolsStore();
  const beadsStore = useBeadsStore();
  const editorStore = useEditorStore();

  const beadCoord = useBeadCoord(options);
  const beadFactory = useBeadFactory();
  const isPainting = ref(false);
  let spanning: ISpanningBead | null = null;

  const paint = createAnimatedFrame((event: MouseEvent, color: string | null) => {
    const point = beadCoord.getFromEvent(event);
    if (!point) return;

    let coord = serializeBeadPoint(point);
    let bead: SchemaBead | null;

    if (spanning) {
      if (spanning.point.x === point.x && spanning.point.y === point.y) {
        return;
      }

      if (spanning.direction) {
        point.x = spanning.direction === 'x' ? point.x : spanning.point.x;
        point.y = spanning.direction === 'y' ? point.y : spanning.point.y;
        coord = serializeBeadPoint(point);
      } else {
        spanning.direction = spanning.point.x === point.x ? 'y' : 'x';
      }

      const existingBead = editorStore.schema.beads[coord];

      if (existingBead
        && isRefBead(existingBead)
        && getBeadSettings(existingBead).to === spanning.coord
      ) {
        return;
      }

      const spanningCoord = serializeBeadPoint(spanning.point);
      bead = beadFactory.createRef(spanningCoord);

      getBeadSettings(spanning.original).span = {
        x: point.x - spanning.point.x,
        y: point.y - spanning.point.y,
      };
    } else {
      const kind = toolsStore.activeBead;
      bead = beadFactory.create(kind, color);

      if (!!bead && isSpannableBead(bead)) {
        spanning = {
          coord,
          point,
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
