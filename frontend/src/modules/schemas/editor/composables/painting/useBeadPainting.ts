import { computed, ref, type Ref } from 'vue';
import { createAnimatedFrame } from '@/helpers';
import { serializeSchemaBeadCoord } from '@/models';
import { PaintEffect, useBeadsStore, useToolsStore } from '../../stores';
import type { IBeadToolsOptions } from './IBeadToolsOptions';
import { useBeadCoord } from './useBeadCoord';
import { useBeadFactory } from './useBeadFactory';

export interface IBeadPaintingListeners {
  mousedown: (event: MouseEvent) => void;
  mousemove?: (event: MouseEvent) => void;
}

export function useBeadPainting(options: IBeadToolsOptions): Ref<IBeadPaintingListeners> {
  const toolsStore = useToolsStore();
  const beadsStore = useBeadsStore();

  const beadCoord = useBeadCoord(options);
  const beadFactory = useBeadFactory();
  const isPainting = ref(false);
  let isSpanning = false;

  const paint = createAnimatedFrame((event: MouseEvent, color: string | null) => {
    const point = beadCoord.getFromEvent(event);
    if (!point) return;

    // TODO: save to bead coords ref of spanning bead

    const coord = serializeSchemaBeadCoord(point.x, point.y);
    const bead = beadFactory.create(toolsStore.activeBead, color);
    const effect = beadsStore.paint(coord, bead);

    if (effect === PaintEffect.EXTENDED) {
      beadCoord.clearCache();
    }
  });

  function onMouseup() {
    isPainting.value = false;
    isSpanning = false;
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
