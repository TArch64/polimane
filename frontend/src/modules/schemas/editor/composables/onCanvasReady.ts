import type { Canvas } from 'fabric';
import { nextTick, onMounted } from 'vue';
import { injectCanvas } from './useCanvas';

export function onCanvasReady(onReady: (canvas: Canvas) => void) {
  const canvas = injectCanvas();

  onMounted(async () => {
    while (!canvas.value) {
      await nextTick();
    }

    onReady(canvas.value);
  });
}
