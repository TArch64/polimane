import type { Canvas, FabricObject } from 'fabric';

export interface IObjectPosition<O extends FabricObject> {
  object: O;
  top: number;
  left: number;
}

export abstract class PositionIterator<O extends FabricObject> extends Iterator<IObjectPosition<O>> {
  readonly [Symbol.iterator] = () => this;

  protected index = 0;

  protected constructor(
    protected readonly canvas: Canvas,
    protected readonly objects: O[],
  ) {
    super();
  }

  protected abstract iteration(): IObjectPosition<O>;

  protected get object(): O {
    return this.objects[this.index];
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
