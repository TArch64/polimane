import { computed, type Ref } from 'vue';
import { useSelectionStore } from '@editor/stores';
import type { IBeadToolsOptions } from './IBeadToolsOptions';

export interface IBeadSelection {
  mousedown: (event: MouseEvent) => void;
  mousemove?: (event: MouseEvent) => void;
}

export function useBeadSelection(options: IBeadToolsOptions): Ref<IBeadSelection> {
  const selectionStore = useSelectionStore();

  function onMouseup() {
    selectionStore.toggleSelecting(false);
  }

  function onMouseDown(event: MouseEvent) {
    selectionStore.toggleSelecting(true);
    selectionStore.setPoint(event.clientX, event.clientY);
    addEventListener('mouseup', onMouseup, { once: true });
  }

  function onMouseMove(event: MouseEvent) {
    selectionStore.extend(event.movementX, event.movementY);
  }

  return computed(() => ({
    mousedown: onMouseDown,
    ...(selectionStore.isSelecting ? { mousemove: onMouseMove } : {}),
  }));
}
