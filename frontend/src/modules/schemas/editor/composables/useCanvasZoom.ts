import { Point, type TPointerEventInfo } from 'fabric';
import { onCanvasReady } from './onCanvasReady';
import { useCanvasCursor } from './useCanvasCursor';

const MIN_ZOOM = 0.5;
const MAX_ZOOM = 10;

export function useCanvasZoom() {
  const cursor = useCanvasCursor();

  onCanvasReady((canvas) => {
    canvas.on('mouse:wheel', (options: TPointerEventInfo<WheelEvent>) => {
      if (!options.e.ctrlKey) {
        return;
      }

      options.e.preventDefault();

      const scaleFactor = 1 - options.e.deltaY / 100;
      const pointer = canvas.getViewportPoint(options.e);
      const point = new Point(pointer.x, pointer.y);

      const zoom = canvas.getZoom();
      const limitedZoom = Math.min(Math.max(zoom * scaleFactor, MIN_ZOOM), MAX_ZOOM);

      canvas.zoomToPoint(point, limitedZoom);

      const affectedObject = canvas.findTarget(options.e);
      cursor.changeTemporarily(zoom > limitedZoom ? 'zoom-out' : 'zoom-in', 100, affectedObject);
    });
  });
}
