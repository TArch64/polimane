import { useCanvasStore, ZOOM_IN_STEP, ZOOM_OUT_STEP } from '@editor/stores';
import { useMouse } from '@vueuse/core';
import type { ICanvasEventListeners } from './useCanvasEvents';
import { useHotKeys } from './useHotKeys';

export function useCanvasZoom(): ICanvasEventListeners {
  const canvasStore = useCanvasStore();
  let lastTouches: Touch[] = [];
  let lastDistance = 0;

  function onWheel(event: WheelEvent): void {
    canvasStore.zoom(event.clientX, event.clientY, event.deltaY);
  }

  function calcDistance(touch1: Touch, touch2: Touch): number {
    return Math.hypot(
      touch2.clientX - touch1.clientX,
      touch2.clientY - touch1.clientY,
    );
  }

  function onTouchMove(event: TouchEvent): void {
    const touch1 = event.touches[0]!;
    const touch2 = event.touches[1]!;

    if (!lastTouches.length) {
      lastTouches = [touch1, touch2];
      lastDistance = calcDistance(touch1, touch2);
      return;
    }

    const distance = calcDistance(touch1, touch2);
    const midPointX = (touch1.clientX + touch2.clientX) / 2;
    const midPointY = (touch1.clientY + touch2.clientY) / 2;

    canvasStore.zoom(midPointX, midPointY, lastDistance - distance);

    lastTouches = [touch1, touch2];
    lastDistance = distance;
  }

  function onTouchEnd() {
    lastTouches = [];
    lastDistance = 0;
  }

  const mouse = useMouse();

  function zoomInPoint(step: number) {
    canvasStore.zoom(mouse.x.value, mouse.y.value, step);
  }

  useHotKeys({
    Meta_Equal: () => zoomInPoint(ZOOM_IN_STEP),
    Meta_Minus: () => zoomInPoint(ZOOM_OUT_STEP),
  });

  return {
    wheel: onWheel,
    touchmove: onTouchMove,
    touchend: onTouchEnd,
  };
}
