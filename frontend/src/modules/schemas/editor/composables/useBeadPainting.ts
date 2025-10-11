import { computed, ref, type Ref } from 'vue';
import { type IPoint, type SchemaBeadCoord, serializeSchemaBeadCoord } from '@/models';
import { createAnimatedFrame } from '@/helpers';
import {
  PaintEffect,
  useBeadsStore,
  useCanvasStore,
  useEditorStore,
  useToolsStore,
} from '../stores';
import { BEAD_CENTER, BEAD_RADIUS, BEAD_SIZE } from './useBeadsGrid';

export interface IBeadPaintingOptions {
  backgroundRectRef: Ref<SVGRectElement>;
}

export interface BeadPainting {
  mousedown: (event: MouseEvent) => void;
  mousemove: (event: MouseEvent) => void;
}

export function useBeadPainting(options: IBeadPaintingOptions) {
  const editorStore = useEditorStore();
  const canvasStore = useCanvasStore();
  const toolsStore = useToolsStore();
  const beadsStore = useBeadsStore();

  const isPainting = ref(false);
  let backgroundRect: DOMRect | null = null;

  function inCircle(mouse: IPoint, circle: IPoint, radius: number) {
    const dx = mouse.x - circle.x;
    const dy = mouse.y - circle.y;
    return dx * dx + dy * dy <= radius * radius;
  }

  function getBeadCoord(event: MouseEvent): SchemaBeadCoord | null {
    const target = event.target as HTMLElement;
    const storedCoord = target.getAttribute('coord');

    if (storedCoord) {
      return storedCoord as SchemaBeadCoord;
    }

    backgroundRect ??= options.backgroundRectRef.value.getBoundingClientRect();

    const mouse: IPoint = {
      y: (event.clientY - backgroundRect.y) / canvasStore.scale,
      x: (event.clientX - backgroundRect.x) / canvasStore.scale,
    };

    const coord: IPoint = {
      y: Math.floor(mouse.y / BEAD_SIZE),
      x: Math.floor(mouse.x / BEAD_SIZE),
    };

    if (coord.y < 0 || coord.x < 0) {
      return null;
    }

    const beadCenter: IPoint = {
      y: (coord.y * BEAD_SIZE) + BEAD_CENTER,
      x: (coord.x * BEAD_SIZE) + BEAD_CENTER,
    };

    if (!inCircle(mouse, beadCenter, BEAD_RADIUS)) {
      return null;
    }

    const { left, top } = editorStore.schema.size;
    return serializeSchemaBeadCoord(coord.x - left, coord.y - top);
  }

  const paint = createAnimatedFrame((event: MouseEvent, color: string | null) => {
    const coord = getBeadCoord(event);
    const effect = coord ? beadsStore.paint(coord, color) : null;

    if (effect === PaintEffect.EXTENDED) {
      backgroundRect = null;
    }
  });

  function onMouseup() {
    isPainting.value = false;
    backgroundRect = null;
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

  return computed((): Partial<BeadPainting> => {
    if (toolsStore.isSelection) {
      return {};
    }

    return {
      mousedown: onMousedown,
      ...(isPainting.value ? { mousemove: onMousemove } : {}),
    };
  });
}
