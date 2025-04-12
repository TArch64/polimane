import type { Ref } from 'vue';
import { Canvas, Point } from 'fabric';
import { useCanvasEvent } from './useCanvasEvent';

const minZoom = 0.1;
const maxZoom = 10;

export function useCanvasZoom(canvas: Ref<Canvas>): void {
  function getZoom(delta: number): number {
    let zoom = canvas.value.getZoom();

    zoom *= 0.999 ** delta;

    if (zoom > maxZoom) return maxZoom;
    if (zoom < minZoom) return minZoom;

    return zoom;
  }

  useCanvasEvent(canvas, 'mouse:wheel', (options) => {
    const { deltaY, offsetX, offsetY } = options.e;
    const zoom = getZoom(deltaY);
    const point = new Point(offsetX, offsetY);

    canvas.value.zoomToPoint(point, zoom);
    options.e.preventDefault();
    options.e.stopPropagation();
  });
}
