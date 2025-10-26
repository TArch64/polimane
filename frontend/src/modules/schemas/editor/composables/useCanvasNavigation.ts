import { useCanvasStore } from '@editor/stores';

export interface ICanvasNavigation {
  navigateWheel: (event: WheelEvent) => void;
  navigateTouch: (event: TouchEvent) => void;
  navigateTouchEnd: () => void;
}

export function useCanvasNavigation(): ICanvasNavigation {
  const canvasStore = useCanvasStore();
  let lastTouch: Touch | null = null;

  function navigateWheel(event: WheelEvent): void {
    const { scale, translation: { x, y } } = canvasStore;

    canvasStore.setTranslation({
      x: x + (event.deltaX / scale),
      y: y + (event.deltaY / scale),
    });
  }

  function navigateTouch(event: TouchEvent): void {
    if (event.touches.length !== 1) return;

    if (!lastTouch) {
      lastTouch = event.touches[0]!;
      return;
    }

    const touch = event.touches[0]!;
    const { scale, translation: { x, y } } = canvasStore;

    canvasStore.setTranslation({
      x: x + ((lastTouch.clientX - touch.clientX) / scale),
      y: y + ((lastTouch.clientY - touch.clientY) / scale),
    });

    lastTouch = touch;
  }

  function navigateTouchEnd() {
    lastTouch = null;
  }

  return { navigateWheel, navigateTouch, navigateTouchEnd };
}
