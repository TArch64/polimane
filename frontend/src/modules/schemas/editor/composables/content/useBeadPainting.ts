import { computed } from 'vue';
import type { SchemaBeadCoord } from '@/models';
import { useBeadsStore, usePaletteStore } from '../../stores';

export function useBeadPainting() {
  const paletteStore = usePaletteStore();
  const beadsStore = useBeadsStore();

  function onMouseup() {
    paletteStore.setPainting(false);
  }

  function paint(event: MouseEvent, color: string | null) {
    const target = event.target as HTMLElement;
    const position = target.getAttribute('coord');

    if (position) {
      beadsStore.paint(position as SchemaBeadCoord, color);
    }
  }

  function onMousedown(event: MouseEvent) {
    if (event.shiftKey) return;

    if (event.buttons === 1) {
      paletteStore.setPainting(true);
      paint(event, paletteStore.activeColor);
    }

    if (event.buttons === 2) {
      paint(event, null);
    }
  }

  function onMousemove(event: MouseEvent) {
    if (event.shiftKey || !paletteStore.isPainting) return;
    paint(event, paletteStore.activeColor);
  }

  return computed(() => ({
    mouseup: onMouseup,
    mousedown: onMousedown,
    ...(paletteStore.isPainting ? { mousemove: onMousemove } : {}),
  }));
}
