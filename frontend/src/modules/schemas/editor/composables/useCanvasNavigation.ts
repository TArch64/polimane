import { useCanvasStore } from '@editor/stores';
import type { IViewBox } from '../types';

export interface ICanvasNavigationOptions {
  viewBox: IViewBox;
}

export interface ICanvasNavigation {
  navigate: (event: WheelEvent) => void;
}

export function useCanvasNavigation(options: ICanvasNavigationOptions): ICanvasNavigation {
  const canvasStore = useCanvasStore();

  function navigate(event: WheelEvent): void {
    options.viewBox.x += (event.deltaX / canvasStore.scale);
    options.viewBox.y += (event.deltaY / canvasStore.scale);
  }

  return { navigate };
}
