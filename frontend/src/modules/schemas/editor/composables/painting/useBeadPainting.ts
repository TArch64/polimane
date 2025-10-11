import { computed, ref, type Ref } from 'vue';
import { createAnimatedFrame } from '@/helpers';
import { PaintEffect, useBeadsStore, useToolsStore } from '../../stores';
import type { IBeadToolsOptions } from './IBeadToolsOptions';
import { useBeadCoord } from './useBeadCoord';

export interface IBeadPainting {
  mousedown: (event: MouseEvent) => void;
  mousemove?: (event: MouseEvent) => void;
}

export function useBeadPainting(options: IBeadToolsOptions): Ref<IBeadPainting> {
  const toolsStore = useToolsStore();
  const beadsStore = useBeadsStore();

  const beadCoord = useBeadCoord(options);
  const isPainting = ref(false);

  const paint = createAnimatedFrame((event: MouseEvent, color: string | null) => {
    const coord = beadCoord.getCoord(event);
    const effect = coord ? beadsStore.paint(coord, color) : null;

    if (effect === PaintEffect.EXTENDED) {
      beadCoord.clearCache();
    }
  });

  function onMouseup() {
    isPainting.value = false;
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
