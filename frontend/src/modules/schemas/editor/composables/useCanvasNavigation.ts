import type { ICanvasZoom } from './useCanvasZoom';

export interface ICanvasNavigation {
  navigate: (event: WheelEvent) => void;
}

export function useCanvasNavigation(canvasZoom: ICanvasZoom): ICanvasNavigation {
  function navigate(event: WheelEvent): void {
    const svg = event.currentTarget as SVGSVGElement;
    const viewBox = svg.viewBox.baseVal;

    viewBox.x += (event.deltaX / canvasZoom.scale);
    viewBox.y += (event.deltaY / canvasZoom.scale);
  }

  return { navigate };
}
