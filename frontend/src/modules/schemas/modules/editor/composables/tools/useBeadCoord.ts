import { useCanvasStore, useEditorStore } from '@editor/stores';
import { BEAD_CIRCLE_CENTER, BEAD_CIRCLE_RADIUS, BEAD_SIZE } from '@editor/const';
import { parseBeadCoord, Point } from '@/models';
import type { IEditorToolOptions } from './tool';

export interface IBeadResolveOptions {
  checkShape?: boolean;
}

export interface IBeadCoord {
  getFromEvent: (event: MouseEvent) => Point | null;
  getFromPoint: (point: Point, options?: IBeadResolveOptions) => Point | null;
  clearCache: () => void;
}

export function useBeadCoord(options: IEditorToolOptions): IBeadCoord {
  const canvasStore = useCanvasStore();
  const editorStore = useEditorStore();

  let backgroundRect: Point | null = null;

  function inCircle(point: Point, circle: Point, radius: number) {
    const dx = point.x - circle.x;
    const dy = point.y - circle.y;
    return dx * dx + dy * dy <= radius * radius;
  }

  function getFromPoint(point: Point, getOptions: IBeadResolveOptions = {}): Point | null {
    if (!backgroundRect) {
      const boundingRect = options.backgroundRef.value.getBoundingClientRect();
      backgroundRect = new Point(boundingRect.x, boundingRect.y);
    }

    const relativePoint = new Point(point.x, point.y)
      .minus(backgroundRect)
      .divide(canvasStore.scale);

    const { left, top } = editorStore.schema.size;
    const y = Math.floor(relativePoint.y / BEAD_SIZE);
    const radialShiftX = options.beadsGrid.getBeadRadialShiftX(y - top);
    const x = Math.floor((relativePoint.x - radialShiftX) / BEAD_SIZE);

    if (y < 0 || x < 0) return null;

    const beadCenter = new Point(x, y)
      .multiply(BEAD_SIZE)
      .plus({ x: radialShiftX + BEAD_CIRCLE_CENTER, y: BEAD_CIRCLE_CENTER });

    if (getOptions.checkShape !== false && !inCircle(relativePoint, beadCenter, BEAD_CIRCLE_RADIUS)) {
      return null;
    }

    return new Point(x - left, y - top);
  }

  function getFromEvent(event: MouseEvent): Point | null {
    const target = event.target as SVGElement;
    const storedCoord = target.getAttribute('coord');

    return storedCoord
      ? parseBeadCoord(storedCoord)
      : getFromPoint(new Point(event.clientX, event.clientY));
  }

  function clearCache() {
    backgroundRect = null;
  }

  return { getFromPoint, getFromEvent, clearCache };
}
