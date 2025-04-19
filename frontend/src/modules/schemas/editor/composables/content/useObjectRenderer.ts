import { computed, type Ref, watch, type WatchStopHandle } from 'vue';
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
  updateObject: (item: I, object: EditorObjectTypeMap[T]) => void;
  updatePositions: (objects: EditorObjectTypeMap[T][]) => void;
}

export function useObjectRenderer<T extends EditorObjectType, I extends IRenderingItem>(options: IObjectRendererOptions<T, I>): void {
  const canvas = injectCanvas();
  const objectRegistry = useObjectRegistry(options.type);
  const stopHandles: Record<string, WatchStopHandle> = {};

  const itemMapping = computed(() => Object.fromEntries(
    options.items.value.map((item) => [item.id, item]),
  ));

  const itemIds = computed(() => options.items.value.map((item) => item.id));

  watch(itemIds, (ids, previousIds = []) => {
    for (const id of ids) {
      const item = itemMapping.value[id];

      if (!previousIds.includes(id)) {
        const object = options.createObject(item);
        objectRegistry.add(id, object);

        stopHandles[id] = watch(() => itemMapping.value[id], (changed) => {
          options.updateObject(changed, object);
          canvas.requestRenderAll();
        }, { deep: true });
      }
    }

    for (const previousId of previousIds) {
      if (!ids.includes(previousId)) {
        objectRegistry.remove(previousId);
        stopHandles[previousId]?.();
      }
    }

    options.updatePositions(ids.map((id) => objectRegistry.get(id)));
    canvas.requestRenderAll();
  }, { immediate: true });
}
