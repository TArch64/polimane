import { onUnmounted } from 'vue';
import { useCanvasCursor } from './useCanvasCursor';
import { onCanvasReady } from './onCanvasReady';

export function useCanvasNavigation() {
  const cursor = useCanvasCursor();
  const abortController = new AbortController();

  onCanvasReady((canvas) => {
    canvas.upperCanvasEl.addEventListener('wheel', (event) => {
      event.preventDefault();

      canvas.viewportTransform[4] -= event.deltaX;
      canvas.viewportTransform[5] -= event.deltaY;
      cursor.changeTemporarily('move', 100);
    }, {
      passive: false,
      signal: abortController.signal,
    });
  });

  onUnmounted(() => {
    abortController.abort();
  });
}
