import { Point } from 'fabric';
import { onCanvasReady } from './onCanvasReady';

const MIN_ZOOM = 0.1;
const MAX_ZOOM = 10;

export const useCanvasZoom = () => onCanvasReady((canvas) => {
  function getZoom(delta: number): number {
    let zoom = canvas.getZoom();

    zoom *= 0.999 ** delta;

    if (zoom > MAX_ZOOM) return MAX_ZOOM;
    if (zoom < MIN_ZOOM) return MIN_ZOOM;

    return zoom;
  }

  canvas.on('mouse:wheel', (options) => {
    const { deltaY, offsetX, offsetY } = options.e;
    const zoom = getZoom(deltaY);
    const point = new Point(offsetX, offsetY);

    canvas.zoomToPoint(point, zoom);
    options.e.preventDefault();
    options.e.stopPropagation();
  });
});
