import { nextTick, onMounted, type Ref } from 'vue';
import type { Canvas, CanvasEvents } from 'fabric';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
type TEventCallback<T = any> = (options: T) => any;

export function useCanvasEvent<
  K extends keyof CanvasEvents,
  E extends CanvasEvents[K],
>(
  canvas: Ref<Canvas>,
  eventName: K,
  handler: TEventCallback<E>,
) {
  async function waitForCanvas(): Promise<void> {
    while (!canvas.value) {
      await nextTick();
    }
  }

  onMounted(async () => {
    await waitForCanvas();
    canvas.value.on(eventName, handler);
  });
}
