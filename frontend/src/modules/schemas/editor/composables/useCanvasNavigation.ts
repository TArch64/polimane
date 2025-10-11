import { useCanvasStore } from '@editor/stores';

export interface ICanvasNavigation {
  navigate: (event: WheelEvent) => void;
}

export function useCanvasNavigation(): ICanvasNavigation {
  const canvasStore = useCanvasStore();

  function navigate(event: WheelEvent): void {
    const { scale, translation: { x, y } } = canvasStore;

    canvasStore.setTranslation({
      x: x + (event.deltaX / scale),
      y: y + (event.deltaY / scale),
    });
  }

  return { navigate };
}
