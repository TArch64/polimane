import { reactive, type Ref } from 'vue';
import { useCanvasStore } from '@editor/stores';
import type { INodeRect } from '@/models';

export interface ICanvasZoomOptions {
  wrapperRect: Ref<DOMRect | null>;
  viewBox: INodeRect;
}

export interface ICanvasZoom {
  zoom: (event: WheelEvent) => void;
}

export function useCanvasZoom(options: ICanvasZoomOptions): ICanvasZoom {
  const canvasStore = useCanvasStore();

  function zoom(event: WheelEvent): void {
    const mousePointToX = (event.clientX / canvasStore.scale) + options.viewBox.x;
    const mousePointToY = (event.clientY / canvasStore.scale) + options.viewBox.y;

    const scaleFactor = 1 - event.deltaY * 0.01;
    canvasStore.setScale(canvasStore.scale * scaleFactor);

    options.viewBox.x = mousePointToX - (event.clientX / canvasStore.scale);
    options.viewBox.y = mousePointToY - (event.clientY / canvasStore.scale);
    options.viewBox.width = options.wrapperRect.value!.width / canvasStore.scale;
    options.viewBox.height = options.wrapperRect.value!.height / canvasStore.scale;
  }

  return reactive({ zoom });
}
