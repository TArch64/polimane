import { onUnmounted } from 'vue';
import type { FabricObject } from 'fabric';
import type { InferObjectEvents } from '@/types';

type EventCallback<P> = (event: P) => void;

export function onObjectEvent<
  O extends FabricObject,
  E extends keyof InferObjectEvents<O>,
>(object: O, name: E, callback: EventCallback<InferObjectEvents<O>[E]>): void {
  // @ts-expect-error hard to set correct type
  const off = object.on(name, callback);

  onUnmounted(() => off());
}

export function onObjectClick(object: FabricObject, callback: () => void): void {
  onObjectEvent(object, 'mousedown', () => {
    object.once('mouseup', () => callback());
  });
}
