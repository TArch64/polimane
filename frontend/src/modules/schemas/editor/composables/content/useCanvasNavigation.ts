import type { ICanvasZoom } from './useCanvasZoom';

export interface ICanvasNavigation {
  navigate: (event: WheelEvent) => void;
}

export function useCanvasNavigation(zoom: ICanvasZoom): ICanvasNavigation {
  function navigate(event: WheelEvent): void {
    const svg = event.currentTarget as SVGSVGElement;
    const viewBox = svg.viewBox.baseVal;

    viewBox.x += (event.deltaX / zoom.scale);
    viewBox.y += (event.deltaY / zoom.scale);
  }

  return { navigate };
}
