import { computed, type Ref, watch } from 'vue';
import { type FabricObject, Rect } from 'fabric';
import type { ISchemaPattern } from '@/models';
import { EditorObjectType, type EditorObjectTypeMap } from '../../enums';
import { injectCanvas } from '../useCanvas';
import { useObjectRegistry } from './useObjectRegistry';
import { COMMON_OBJECT_PROPS } from './commonObjectProps';

const CANVAS_PADDING = 10;

export function usePatternRenderer(patterns: Ref<ISchemaPattern[]>) {
  const canvas = injectCanvas();
  const objectRegistry = useObjectRegistry();
  const ids = computed(() => patterns.value.map((pattern) => pattern.id));

  function createObject(): EditorObjectTypeMap[EditorObjectType.PATTERN] {
    return new Rect({
      ...COMMON_OBJECT_PROPS,
      left: CANVAS_PADDING,
      top: CANVAS_PADDING,
      width: 100,
      height: 10,
      fill: 'red',
    });
  }

  function updatePositions(objects: FabricObject[]): void {
    const freeSpaceX = canvas.value.width - CANVAS_PADDING * 2;
    const totalHeight = objects.reduce((acc, object, index) => acc + object.height + (index < objects.length ? 10 : 0), 0);
    let nextOffsetY = (canvas.value.height - totalHeight - CANVAS_PADDING * 2) / 2;

    for (const object of objects) {
      if (object.top !== nextOffsetY) {
        object.setY(nextOffsetY);
      }

      const offsetLeft = (freeSpaceX - object.width) / 2;

      if (object.left !== offsetLeft) {
        object.setX(offsetLeft);
      }

      nextOffsetY += object.height + 10;
    }
  }

  watch(ids, (ids, previousIds = []) => {
    for (const id of ids) {
      if (!previousIds.includes(id)) {
        objectRegistry.pattern.add(id, createObject());
      }
    }

    for (const previousId of previousIds) {
      if (!ids.includes(previousId)) {
        objectRegistry.pattern.remove(previousId);
      }
    }

    updatePositions(ids.map((id) => objectRegistry.pattern.get(id)));
    canvas.value.requestRenderAll();
  }, { immediate: true });
}
