import type { IViewBox } from '../types';
import type { ICanvasZoom } from './useCanvasZoom';

export interface ICanvasNavigationOptions {
  canvasZoom: ICanvasZoom;
  viewBox: IViewBox;
}

export interface ICanvasNavigation {
  navigate: (event: WheelEvent) => void;
}

export function useCanvasNavigation(options: ICanvasNavigationOptions): ICanvasNavigation {
  function navigate(event: WheelEvent): void {
    options.viewBox.x += (event.deltaX / options.canvasZoom.scale);
    options.viewBox.y += (event.deltaY / options.canvasZoom.scale);
  }

  return { navigate };
}
