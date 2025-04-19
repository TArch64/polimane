import type { Canvas } from 'fabric';
import { nextTick, onMounted } from 'vue';
import { injectCanvasRef } from './useCanvas';

export function onCanvasReady(onReady: (canvas: Canvas) => void) {
  const canvas = injectCanvasRef();

  onMounted(async () => {
    while (!canvas.value) {
      await nextTick();
    }

    onReady(canvas.value);
  });
}
