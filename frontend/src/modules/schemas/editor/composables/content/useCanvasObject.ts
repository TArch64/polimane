import { markRaw, onUnmounted } from 'vue';
import type { FabricObject } from 'fabric';
import type { ISchemaObject } from '@/models';
import { watchObject } from './watchObject';
import { useObjectParent } from './useObjectParent';
import { isObjectImplementsOnUpdate } from './IObjectOnUpdate';
import { isObjectImplementsOnAdded } from './IObjectOnAdded';

const objects = new Map<string, FabricObject>();

export function useCanvasObject<O extends FabricObject>(id: string, create: () => O): O {
  const parent = useObjectParent();
  const object = markRaw(create());

  objects.set(id, object);
  parent.add(object);

  if (isObjectImplementsOnAdded(object)) {
    object.onAdded(parent);
  }

  onUnmounted(() => {
    objects.delete(id);
    parent.remove(object);
    object.dispose();
  });

  return object;
}

export function useCanvasEntityObject<
  D extends ISchemaObject,
  O extends FabricObject,
>(data: D, create: () => O): O {
  const object = useCanvasObject(data.id, create);

  if (isObjectImplementsOnUpdate(object)) {
    watchObject(() => data, () => {
      object.onUpdate(data);
      object.canvas?.requestRenderAll();
    });
  }

  return object;
}

export function getCanvasObject<O extends FabricObject>(id: string): O | null {
  return objects.get(id) as O ?? null;
}
