import { reactive, type Ref, ref } from 'vue';
import type { IViewBox } from '../types';

const MIN_ZOOM = 0.5;
const MAX_ZOOM = 10;

export interface ICanvasZoomOptions {
  wrapperRect: Ref<DOMRect | null>;
  viewBox: IViewBox;
}

export interface ICanvasZoom {
  scale: number;
  zoom: (event: WheelEvent) => void;
}

export function useCanvasZoom(options: ICanvasZoomOptions): ICanvasZoom {
  const scale = ref(1);

  function zoom(event: WheelEvent): void {
    const mousePointToX = (event.clientX / scale.value) + options.viewBox.x;
    const mousePointToY = (event.clientY / scale.value) + options.viewBox.y;

    const scaleFactor = 1 - event.deltaY * 0.01;
    scale.value = Math.min(Math.max(scale.value * scaleFactor, MIN_ZOOM), MAX_ZOOM);

    options.viewBox.x = mousePointToX - (event.clientX / scale.value);
    options.viewBox.y = mousePointToY - (event.clientY / scale.value);
    options.viewBox.width = options.wrapperRect.value!.width / scale.value;
    options.viewBox.height = options.wrapperRect.value!.height / scale.value;
  }

  return reactive({ zoom, scale });
}
