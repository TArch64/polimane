import Konva from 'konva';
import { computed, type Ref, watch } from 'vue';
import { useCursorStore, useDraggingStore } from '../../stores';

type Cursor = 'default' | 'crosshair' | 'grab' | 'grabbing';

export function useCanvasCursor(stage: Ref<Konva.Stage>) {
  const cursorStore = useCursorStore();
  const draggingStore = useDraggingStore();

  const cursor = computed((): Cursor => {
    if (cursorStore.isDragging) {
      return draggingStore.draggingObject ? 'grabbing' : 'grab';
    }
    if (cursorStore.isPainting) {
      return 'crosshair';
    }
    return 'default';
  });

  watch(cursor, (cursor) => {
    stage.value.container().style.cursor = cursor;
  });
}
