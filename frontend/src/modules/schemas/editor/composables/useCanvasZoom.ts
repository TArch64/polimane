import { reactive } from 'vue';
import { useCanvasStore } from '@editor/stores';

export interface ICanvasZoom {
  zoom: (event: WheelEvent) => void;
}

export function useCanvasZoom(): ICanvasZoom {
  const canvasStore = useCanvasStore();

  function zoom(event: WheelEvent): void {
    let scale = canvasStore.scale;
    const mousePointToX = (event.clientX / scale) + canvasStore.translation.x;
    const mousePointToY = (event.clientY / scale) + canvasStore.translation.y;

    const scaleFactor = 1 - event.deltaY * 0.01;
    scale = canvasStore.setScale(scale * scaleFactor);

    canvasStore.setTranslation({
      x: mousePointToX - (event.clientX / scale),
      y: mousePointToY - (event.clientY / scale),
    });
  }

  return reactive({ zoom });
}
