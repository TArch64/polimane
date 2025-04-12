import { Point, type TPointerEventInfo } from 'fabric';
import type { BrowserCursor } from '@/types';
import { onCanvasReady } from './onCanvasReady';

function pointFromEvent(options: TPointerEventInfo): Point {
  const event = options.e as MouseEvent;
  return new Point(event.clientX, event.clientY);
}

export const useCanvasNavigation = () => onCanvasReady((canvas) => {
  let lastPosition = new Point();

  function onMouseMove(options: TPointerEventInfo): void {
    const newPosition = pointFromEvent(options);

    canvas.viewportTransform[4] += newPosition.x - lastPosition.x;
    canvas.viewportTransform[5] += newPosition.y - lastPosition.y;

    lastPosition = newPosition;
    canvas.requestRenderAll();
  }

  function setCursor(options: TPointerEventInfo, cursor: BrowserCursor) {
    canvas.defaultCursor = cursor;
    canvas.setCursor(cursor);

    const affectedObject = canvas.findTarget(options.e);

    if (affectedObject) {
      affectedObject.hoverCursor = cursor;
    }

    canvas.requestRenderAll();
  }

  canvas.on('mouse:down', (downOptions) => {
    if (!(downOptions.e as MouseEvent).altKey) {
      return;
    }

    setCursor(downOptions, 'move');
    lastPosition = pointFromEvent(downOptions);

    const unsubscribe = canvas.on('mouse:move', onMouseMove);

    canvas.once('mouse:up', () => {
      unsubscribe();
      lastPosition = new Point();
      setCursor(downOptions, 'default');
    });
  });
});
