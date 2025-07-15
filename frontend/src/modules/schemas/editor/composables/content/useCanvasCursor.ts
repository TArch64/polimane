import { computed, type Ref, watch } from 'vue';
import Konva from 'konva';
import { usePaletteStore } from '@/modules/schemas/editor/stores';

type Cursor = 'default' | 'crosshair';

export function useCanvasCursor(stage: Ref<Konva.Stage>) {
  const paletteStore = usePaletteStore();

  const cursor = computed((): Cursor => {
    if (paletteStore.isPainting) {
      return 'crosshair';
    }
    return 'default';
  });

  watch(cursor, (cursor) => {
    stage.value.container().style.cursor = cursor;
  });
}
