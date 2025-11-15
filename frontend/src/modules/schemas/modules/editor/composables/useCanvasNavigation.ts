import { useCanvasStore } from '@editor/stores';
import type { ICanvasEventListeners } from './useCanvasEvents';

export function useCanvasNavigation(): ICanvasEventListeners {
  const canvasStore = useCanvasStore();
  let lastTouch: Touch | null = null;

  function onWheel(event: WheelEvent): void {
    canvasStore.navigate(event.deltaX, event.deltaY);
  }

  function onTouchMove(event: TouchEvent): void {
    if (!lastTouch) {
      lastTouch = event.touches[0]!;
      return;
    }

    const touch = event.touches[0]!;
    canvasStore.navigate(lastTouch.clientX - touch.clientX, lastTouch.clientY - touch.clientY);
    lastTouch = touch;
  }

  function onTouchEnd() {
    lastTouch = null;
  }

  return {
    wheel: onWheel,
    touchmove: onTouchMove,
    touchend: onTouchEnd,
  };
}
