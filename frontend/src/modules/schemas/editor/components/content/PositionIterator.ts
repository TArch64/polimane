import type { FabricObject } from 'fabric';
import type { ObjectParent } from '@/modules/schemas/editor/composables';

export interface IObjectPosition<O extends FabricObject> {
  object: O;
  top: number;
  left: number;
}

export abstract class PositionIterator<O extends FabricObject> extends Iterator<IObjectPosition<O>> {
  readonly [Symbol.iterator] = () => this;

  protected index = 0;

  protected constructor(
    protected readonly parent: ObjectParent,
    protected readonly objects: O[],
  ) {
    super();
  }

  protected abstract iteration(): IObjectPosition<O>;

  protected get object(): O {
    return this.objects[this.index];
  }

  protected calcListSize(side: 'width' | 'height', gap: number): number {
    return this.objects.reduce((acc, object, index) => {
      const offset = index < this.objects.length ? gap : 0;
      return acc + object[side] + offset;
    }, 0);
  }

  next(): IteratorResult<IObjectPosition<O>> {
    if (!this.objects[this.index]) {
      return { done: true, value: undefined };
    }

    const value = this.iteration();
    this.index++;

    return { done: false, value };
  }
}
