import { computed, ref, type Ref } from 'vue';
import { createAnimatedFrame } from '@/helpers';
import {
  getBeadSettings,
  type IPoint,
  type SchemaBead,
  type SchemaSpannableBead,
  serializeBeadCoord,
} from '@/models';
import { isBeadSpannableKind } from '@/enums';
import { PaintEffect, useBeadsStore, useToolsStore } from '../../stores';
import type { IBeadToolsOptions } from './IBeadToolsOptions';
import { useBeadCoord } from './useBeadCoord';
import { useBeadFactory } from './useBeadFactory';

export interface IBeadPaintingListeners {
  mousedown: (event: MouseEvent) => void;
  mousemove?: (event: MouseEvent) => void;
}

interface ISpanningBead {
  coord: IPoint;
  original: SchemaSpannableBead;
}

export function useBeadPainting(options: IBeadToolsOptions): Ref<IBeadPaintingListeners> {
  const toolsStore = useToolsStore();
  const beadsStore = useBeadsStore();

  const beadCoord = useBeadCoord(options);
  const beadFactory = useBeadFactory();
  const isPainting = ref(false);
  let spanning: ISpanningBead | null = null;

  const paint = createAnimatedFrame((event: MouseEvent, color: string | null) => {
    const point = beadCoord.getFromEvent(event);
    if (!point) return;

    const coord = serializeBeadCoord(point.x, point.y);
    let bead: SchemaBead | null;

    if (spanning) {
      if (spanning.coord.x === point.x && spanning.coord.y === point.y) {
        return;
      }

      const spanningCoord = serializeBeadCoord(spanning.coord.x, spanning.coord.y);
      bead = beadFactory.createRef(spanningCoord);

      getBeadSettings(spanning.original).span = {
        x: point.x - spanning.coord.x,
        y: point.y - spanning.coord.y,
      };
    } else {
      const kind = toolsStore.activeBead;
      bead = beadFactory.create(kind, color);

      if (!!bead && isBeadSpannableKind(kind)) {
        spanning = {
          coord: point,
          original: bead as SchemaSpannableBead,
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
