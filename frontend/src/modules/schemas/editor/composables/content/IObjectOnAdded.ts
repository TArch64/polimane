import type { FabricObject } from 'fabric';
import type { ObjectParent } from './useObjectParent';

export interface IObjectOnAdded {
  onAdded(parent: ObjectParent): void;
}

export function isObjectImplementsOnAdded<
  O extends FabricObject,
>(object: O): object is O & IObjectOnAdded {
  return 'onAdded' in object;
}
