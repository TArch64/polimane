import { computed, type Ref, watch } from 'vue';
import Konva from 'konva';
import { useCursorStore } from '@/modules/schemas/editor/stores';

type Cursor = 'default' | 'crosshair' | 'grab';

export function useCanvasCursor(stage: Ref<Konva.Stage>) {
  const cursorStore = useCursorStore();

  const cursor = computed((): Cursor => {
    if (cursorStore.isDragging) {
      return 'grab';
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
