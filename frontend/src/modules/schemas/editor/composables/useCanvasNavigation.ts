import type { Ref } from 'vue';
import type { Canvas, TPointerEventInfo } from 'fabric';
import { useCanvasEvent } from './useCanvasEvent';

export function useCanvasNavigation(canvas: Ref<Canvas>) {
  let lastPosX: number, lastPosY: number;

  function onMouseMove(options: TPointerEventInfo<MouseEvent>): void {
    const event = options.e;
    const deltaX = event.clientX - lastPosX;
    const deltaY = event.clientY - lastPosY;

    const vpt = canvas.value.viewportTransform;
    vpt[4] += deltaX;
    vpt[5] += deltaY;

    canvas.value.requestRenderAll();
    lastPosX = event.clientX;
    lastPosY = event.clientY;
  }

  useCanvasEvent(canvas, 'mouse:down', (options) => {
    const event = options.e as MouseEvent;
    lastPosX = event.clientX;
    lastPosY = event.clientY;

    canvas.value.on('mouse:move', onMouseMove);

    canvas.value.once('mouse:up', () => {
      canvas.value.off('mouse:move', onMouseMove);
      lastPosX = 0;
      lastPosY = 0;
    });
  });
}
