import { onUnmounted } from 'vue';
import { onCanvasReady } from './onCanvasReady';

export function useCanvasNavigation() {
  const abortController = new AbortController();

  onCanvasReady((canvas) => {
    canvas.upperCanvasEl.addEventListener('wheel', (event) => {
      if (event.ctrlKey) {
        return;
      }

      event.preventDefault();
      canvas.viewportTransform[4] -= event.deltaX;
      canvas.viewportTransform[5] -= event.deltaY;
      canvas.forEachObject((object) => object.setCoords());
      canvas.requestRenderAll();
    }, {
      passive: false,
      signal: abortController.signal,
    });
  });

  onUnmounted(() => {
    abortController.abort();
  });
}
