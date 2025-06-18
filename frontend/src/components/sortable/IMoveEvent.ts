import type { ISortableEntity } from './ISortableEntity';

export interface IMoveEvent<I extends ISortableEntity> {
  fromIndex: number;
  toIndex: number;
  item: I;
  updated: I[];
}
