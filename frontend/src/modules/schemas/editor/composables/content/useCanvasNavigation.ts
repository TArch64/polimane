import type { KonvaEventObject } from 'konva/lib/Node';
import Konva from 'konva';

export interface ICanvasNavigation {
  navigate: (kEvent: KonvaEventObject<WheelEvent, Konva.Stage>) => void;
}

export function useCanvasNavigation(): ICanvasNavigation {
  function navigate(kEvent: KonvaEventObject<WheelEvent, Konva.Stage>): void {
    const { evt: event, currentTarget: stage } = kEvent;
    const currentPos = stage.position();

    stage.position({
      x: currentPos.x - event.deltaX,
      y: currentPos.y - event.deltaY,
    });
  }

  return { navigate };
}
