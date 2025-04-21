import { markRaw, onUnmounted } from 'vue';
import type { FabricObject } from 'fabric';
import type { ISchemaObject } from '@/models';
import { watchObject } from '@/modules/schemas/editor/composables';
import { injectCanvas } from '../useCanvas';

const objects = new Map<string, FabricObject>();

export interface IUpdatableFabricObject<D> {
  update(data: D): void;
}

export function useCanvasObject<
  D extends ISchemaObject,
  O extends FabricObject & IUpdatableFabricObject<D>,
>(data: D, create: () => O): O {
  const canvas = injectCanvas();

  const object = markRaw(create());

  objects.set(data.id, object);
  canvas.add(object);

  watchObject(() => data, (patch) => object.update(patch as D));

  onUnmounted(() => {
    objects.delete(data.id);
    canvas.remove(object);
  });

  return object;
}

export function getCanvasObject<O extends FabricObject>(id: string): O | null {
  return objects.get(id) as O ?? null;
}
