import { useCanvasStore, useEditorStore } from '@editor/stores';
import { type IPoint, type SchemaBeadCoord, serializeSchemaBeadCoord } from '@/models';
import { BEAD_CENTER, BEAD_RADIUS, BEAD_SIZE } from '../useBeadsGrid';
import type { IBeadToolsOptions } from './IBeadToolsOptions';

export interface IBeadCoord {
  getCoord: (event: MouseEvent) => SchemaBeadCoord | null;
  clearCache: () => void;
}

export function useBeadCoord(options: IBeadToolsOptions): IBeadCoord {
  const canvasStore = useCanvasStore();
  const editorStore = useEditorStore();

  let backgroundRect: DOMRect | null = null;

  function inCircle(mouse: IPoint, circle: IPoint, radius: number) {
    const dx = mouse.x - circle.x;
    const dy = mouse.y - circle.y;
    return dx * dx + dy * dy <= radius * radius;
  }

  function getCoord(event: MouseEvent): SchemaBeadCoord | null {
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

  function clearCache() {
    backgroundRect = null;
  }

  return { getCoord, clearCache };
}
