import { toRaw } from 'vue';
import type { ISchemaObject } from './ISchemaObject';
import type { InferSchemaContent, ISchemaWithContent } from './ISchemaWithContent';

export interface ICollectionOptions<P extends ISchemaWithContent, O extends ISchemaObject> {
  onAdded?: (parent: P, object: O) => void;
}

export interface ICollectionInsertOptions {
  toIndex?: number;
}

export class Collection<P extends ISchemaWithContent, O extends ISchemaObject = InferSchemaContent<P>> {
  private static cache = new WeakMap<ISchemaWithContent, Collection<ISchemaWithContent>>();

  static fromParent<
    P extends ISchemaWithContent,
    O extends ISchemaObject = InferSchemaContent<P>,
  >(parent: P, options?: ICollectionOptions<P, O>): Collection<P, O> {
    const cached = Collection.cache.get(toRaw(parent));

    if (cached) {
      return cached as unknown as Collection<P, O>;
    }

    const collection = new Collection<P, O>(parent, options);
    Collection.cache.set(toRaw(parent), collection as unknown as Collection<ISchemaWithContent>);
    return collection;
  }

  private constructor(
    private readonly parent: P,
    private readonly options: ICollectionOptions<P, O> = {},
  ) {
  }

  get values(): O[] {
    return this.parent.content as O[];
  }

  set values(values: O[]) {
    this.parent.content = values;
  }

  get first(): O | null {
    return this.values[0] ?? null;
  }

  get size(): number {
    return this.values.length;
  }

  insert(item: O, options?: ICollectionInsertOptions): O;
  insert(list: O[], options?: ICollectionInsertOptions): O[];
  insert(singleOrList: O | O[], options: ICollectionInsertOptions = {}): O | O[] {
    const list = Array.isArray(singleOrList) ? singleOrList : [singleOrList];
    const toIndex = options.toIndex ?? -1;
    toIndex === -1 ? this.values.push(...list) : this.values.splice(toIndex, 0, ...list);

    if (this.options.onAdded) {
      for (const item of list) {
        this.options.onAdded(this.parent, item);
      }
    }

    return singleOrList;
  }

  delete(item: O): void {
    const index = this.values.findIndex((i) => i.id === item.id);
    this.values.splice(index, 1);
  }

  indexOf(item: O): number {
    return this.values.findIndex((i) => i.id === item.id);
  }
}
