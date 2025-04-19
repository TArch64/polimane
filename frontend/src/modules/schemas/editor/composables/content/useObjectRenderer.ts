import { computed, type Ref, watch } from 'vue';
import { EditorObjectType } from '../../enums';
import { injectCanvas } from '../useCanvas';
import { useObjectRegistry } from './useObjectRegistry';
import type { EditorObjectTypeMap } from './objects';

export interface IRenderingItem {
  id: string;
}

export interface IObjectRendererOptions<T extends EditorObjectType, I extends IRenderingItem> {
  type: T;
  items: Ref<I[]>;
  createObject: (item: I) => EditorObjectTypeMap[T];
  updatePositions: (objects: EditorObjectTypeMap[T][]) => void;
}

export function useObjectRenderer<T extends EditorObjectType, I extends IRenderingItem>(options: IObjectRendererOptions<T, I>): void {
  const canvas = injectCanvas();
  const objectRegistry = useObjectRegistry();

  const itemMapping = computed(() => Object.fromEntries(
    options.items.value.map((item) => [item.id, item]),
  ));

  const itemIds = computed(() => options.items.value.map((item) => item.id));

  watch(itemIds, (ids, previousIds = []) => {
    for (const id of ids) {
      const item = itemMapping.value[id];

      if (!previousIds.includes(id)) {
        objectRegistry.pattern.add(id, options.createObject(item));
      }
    }

    for (const previousId of previousIds) {
      if (!ids.includes(previousId)) {
        objectRegistry.pattern.remove(previousId);
      }
    }

    options.updatePositions(ids.map((id) => objectRegistry.pattern.get(id)));
    canvas.requestRenderAll();
  }, { immediate: true });
}
