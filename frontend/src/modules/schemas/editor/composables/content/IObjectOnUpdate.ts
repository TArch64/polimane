import type { FabricObject } from 'fabric';
import type { ISchemaObject } from '@/models';

export interface IObjectOnUpdate<D> {
  onUpdate(data: D): void;
}

export function isObjectImplementsOnUpdate<
  D extends ISchemaObject,
  O extends FabricObject,
>(object: O): object is O & IObjectOnUpdate<D> {
  return 'onUpdate' in object;
}
