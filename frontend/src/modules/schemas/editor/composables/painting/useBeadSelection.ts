import { computed, ref, type Ref } from 'vue';
import { useSelectionStore } from '@editor/stores';
import type { IBeadToolsOptions } from './IBeadToolsOptions';

export interface IBeadSelection {
  mousedown: (event: MouseEvent) => void;
}

export function useBeadSelection(options: IBeadToolsOptions): Ref<IBeadSelection> {
  const selectionStore = useSelectionStore();

  const isSelecting = ref(false);
  let backgroundRect: DOMRect | null = null;

  function onMouseup() {
    isSelecting.value = false;
    backgroundRect = null;
    selectionStore.reset();
  }

  function onMouseDown(event: MouseEvent) {
    backgroundRect ??= options.backgroundRectRef.value.getBoundingClientRect();

    isSelecting.value = true;
    selectionStore.setPoint(event.clientX, event.clientY);

    addEventListener('mouseup', onMouseup, { once: true });
  }

  function onMouseMove(event: MouseEvent) {
    selectionStore.extend(event.movementX, event.movementY);
  }

  return computed(() => ({
    mousedown: onMouseDown,
    ...(isSelecting.value ? { mousemove: onMouseMove } : {}),
  }));
}
