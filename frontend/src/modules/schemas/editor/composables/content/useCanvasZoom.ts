import type { KonvaEventObject } from 'konva/lib/Node';
import Konva from 'konva';

const MIN_ZOOM = 0.5;
const MAX_ZOOM = 10;

export interface ICanvasZoom {
  zoom: (kEvent: KonvaEventObject<WheelEvent, Konva.Stage>) => void;
}

export function useCanvasZoom(): ICanvasZoom {
  function zoom(kEvent: KonvaEventObject<WheelEvent, Konva.Stage>): void {
    const { evt: event, currentTarget: stage } = kEvent;

    const oldScale = stage.scaleX();
    const pointer = stage.getPointerPosition()!;

    const mousePointTo = {
      x: (pointer.x - stage.x()) / oldScale,
      y: (pointer.y - stage.y()) / oldScale,
    };

    const scaleFactor = 1 - event.deltaY * 0.01;
    const newScale = Math.min(Math.max(oldScale * scaleFactor, MIN_ZOOM), MAX_ZOOM);

    stage.scale({
      x: newScale,
      y: newScale,
    });

    const newPos = {
      x: pointer.x - mousePointTo.x * newScale,
      y: pointer.y - mousePointTo.y * newScale,
    };

    stage.position(newPos);
  }

  return { zoom };
}
