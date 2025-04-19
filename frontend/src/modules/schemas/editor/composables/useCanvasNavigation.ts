import { Point, type TPointerEventInfo } from 'fabric';
import { useCanvasCursor } from './useCanvasCursor';
import { onCanvasReady } from './onCanvasReady';

function pointFromEvent(options: TPointerEventInfo): Point {
  const event = options.e as MouseEvent;
  return new Point(event.clientX, event.clientY);
}

export function useCanvasNavigation() {
  const cursor = useCanvasCursor();

  onCanvasReady((canvas) => {
    let lastPosition = new Point();

    function onMouseMove(options: TPointerEventInfo): void {
      const newPosition = pointFromEvent(options);

      canvas.viewportTransform[4] += newPosition.x - lastPosition.x;
      canvas.viewportTransform[5] += newPosition.y - lastPosition.y;

      lastPosition = newPosition;
      canvas.requestRenderAll();
    }

    canvas.on('mouse:down', (downOptions) => {
      const event = downOptions.e as MouseEvent;

      if (event.buttons !== 2) {
        return;
      }

      const affectedObject = canvas.findTarget(downOptions.e);
      cursor.change('move', affectedObject);

      lastPosition = pointFromEvent(downOptions);

      const unsubscribe = canvas.on('mouse:move', onMouseMove);

      canvas.once('mouse:up', () => {
        unsubscribe();
        lastPosition = new Point();
        cursor.change('default', affectedObject);
      });
    });
  });
}
