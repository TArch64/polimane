import { useCanvasZoom } from './useCanvasZoom';
import { useCanvasNavigation } from './useCanvasNavigation';

export interface ICanvasEventListeners {
  wheel: (event: WheelEvent) => void;
  touchmove: (event: TouchEvent) => void;
  touchend: () => void;
}

export interface ICanvasEvents {
  listeners: ICanvasEventListeners;
}

export function useCanvasEvents(): ICanvasEvents {
  const canvasZoom = useCanvasZoom();
  const canvasNavigation = useCanvasNavigation();
  let lastEvent: TouchEvent | null = null;

  function onWheel(event: WheelEvent) {
    event.preventDefault();

    event.ctrlKey
      ? canvasZoom.wheel(event)
      : canvasNavigation.wheel(event);
  }

  const touchHandlers: Record<number, ICanvasEventListeners> = {
    1: canvasNavigation,
    2: canvasZoom,
  };

  function onTouchMove(event: TouchEvent) {
    event.preventDefault();
    touchHandlers[event.touches.length]?.touchmove(event);
    lastEvent = event;
  }

  function onTouchEnd() {
    if (!lastEvent) return;
    touchHandlers[lastEvent.touches.length]?.touchend?.();
    lastEvent = null;
  }

  return {
    listeners: {
      wheel: onWheel,
      touchmove: onTouchMove,
      touchend: onTouchEnd,
    },
  };
}
