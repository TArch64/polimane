import { defineStore } from 'pinia';
import { onScopeDispose, type Ref, ref } from 'vue';
import type { ISchemaObject } from '@/models';

export type OnDragCleanUp = () => void;

export interface IDragCaptureOptions {
  object: ISchemaObject;
  dragTranslation: Ref<number>;
  onCleanUp: () => void;
}

export const useDraggingStore = defineStore('schemas/editor/dragging', () => {
  const draggingObject = ref<ISchemaObject | null>(null);
  let dragTranslation = ref(0);
  let abortController: AbortController | null = null;
  let onCleanUp: OnDragCleanUp | null = null;

  function release(): void {
    abortController?.abort();
    abortController = null;
    draggingObject.value = null;
    dragTranslation.value = 0;
    dragTranslation = ref(0);
    onCleanUp?.();
  }

  function dragObject(event: MouseEvent) {
    dragTranslation.value += event.movementY;
  }

  function capture(options: IDragCaptureOptions): void {
    draggingObject.value = options.object;
    onCleanUp = options.onCleanUp;
    dragTranslation = options.dragTranslation;
    abortController = new AbortController();
    addEventListener('mouseup', release, { signal: abortController.signal });
    addEventListener('mousemove', dragObject, { signal: abortController.signal });
  }

  onScopeDispose(release);

  return {
    capture,
    draggingObject,
  };
});
