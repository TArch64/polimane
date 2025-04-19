import { type Ref, toRef } from 'vue';

export interface ICollectionItem {
  id: string;
}

type CollectionKeys<T extends object> = keyof {
  [K in keyof T as T[K] extends ICollectionItem[] ? K : never]: T[K];
};

export class Collection<I extends ICollectionItem> {
  static fromProperty<T extends object, K extends CollectionKeys<T>>(object: T, key: K): Collection<Extract<T[K], ICollectionItem[]>[number]> {
    return Collection.fromRef(toRef(object, key) as Ref<Extract<T[K], ICollectionItem[]>>);
  }

  static fromRef<I extends ICollectionItem>(ref: Ref<I[]>): Collection<I> {
    return new Collection(() => ref.value, (items) => ref.value = items);
  }

  constructor(
    private readonly getSource: () => I[],
    private readonly setSource: (items: I[]) => void,
  ) {
  }

  get values(): I[] {
    return this.getSource();
  }

  get first(): I | null {
    return this.values[0] ?? null;
  }

  get size(): number {
    return this.values.length;
  }

  append(item: I): void {
    this.values.push(item);
    return item;
  }

  delete(item: I): void {
    const index = this.values.findIndex((i) => i.id === item.id);
    this.values.splice(index, 1);
  }
}
