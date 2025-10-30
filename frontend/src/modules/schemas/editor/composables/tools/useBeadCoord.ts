import { useCanvasStore, useEditorStore } from '@editor/stores';
import { BEAD_CIRCLE_CENTER, BEAD_CIRCLE_RADIUS, BEAD_SIZE } from '@editor/const';
import { type IPoint, parseBeadCoord } from '@/models';
import type { IEditorToolOptions } from './tool';

export interface IBeadResolveOptions {
  checkShape?: boolean;
}

export interface IBeadCoord {
  getFromEvent: (event: MouseEvent) => IPoint | null;
  getFromPoint: (point: IPoint, options?: IBeadResolveOptions) => IPoint | null;
  clearCache: () => void;
}

export function useBeadCoord(options: IEditorToolOptions): IBeadCoord {
  const canvasStore = useCanvasStore();
  const editorStore = useEditorStore();

  let backgroundRect: DOMRect | null = null;

  function inCircle(point: IPoint, circle: IPoint, radius: number) {
    const dx = point.x - circle.x;
    const dy = point.y - circle.y;
    return dx * dx + dy * dy <= radius * radius;
  }

  function getFromPoint(point: IPoint, getOptions: IBeadResolveOptions = {}): IPoint | null {
    backgroundRect ??= options.backgroundRef.value.getBoundingClientRect();

    const relativePoint: IPoint = {
      y: (point.y - backgroundRect.y) / canvasStore.scale,
      x: (point.x - backgroundRect.x) / canvasStore.scale,
    };

    const coord: IPoint = {
      y: Math.floor(relativePoint.y / BEAD_SIZE),
      x: Math.floor(relativePoint.x / BEAD_SIZE),
    };

    if (coord.y < 0 || coord.x < 0) {
      return null;
    }

    const beadCenter: IPoint = {
      y: (coord.y * BEAD_SIZE) + BEAD_CIRCLE_CENTER,
      x: (coord.x * BEAD_SIZE) + BEAD_CIRCLE_CENTER,
    };

    if (getOptions.checkShape !== false && !inCircle(relativePoint, beadCenter, BEAD_CIRCLE_RADIUS)) {
      return null;
    }

    const { left, top } = editorStore.schema.size;
    return { x: coord.x - left, y: coord.y - top };
  }

  function getFromEvent(event: MouseEvent): IPoint | null {
    const target = event.target as SVGElement;
    const storedCoord = target.getAttribute('coord');

    return storedCoord
      ? parseBeadCoord(storedCoord)
      : getFromPoint({ x: event.clientX, y: event.clientY });
  }

  function clearCache() {
    backgroundRect = null;
  }

  return { getFromPoint, getFromEvent, clearCache };
}
