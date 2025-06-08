import { defineStore, getActivePinia, type Store } from 'pinia';
import type { SafeAny } from '@/types';
import { destroyStore, fetchMapValue } from '@/helpers';
import type { InferStore } from './InferStore';

export interface IDynamicStoreOptions<A extends SafeAny[] = SafeAny[]> {
  $args?: A;
  buildPath: (...value: A) => string;
  setup: (...value: A) => SafeAny;
}

type InferStoreArgs<O extends IDynamicStoreOptions> = Exclude<O['$args'], undefined>;
type InferStoreFromOptions<O extends IDynamicStoreOptions> = InferStore<ReturnType<O['buildPath']>, O['setup']>;

export class DynamicStore<O extends IDynamicStoreOptions> {
  private stores = new Map<string, InferStoreFromOptions<O>>();
  readonly $storeType!: InferStoreFromOptions<O>;

  constructor(private readonly options: O) {
  }

  build() {
    return {
      useStore: this.useStore.bind(this),
      disposeStores: this.disposeStores.bind(this),
    };
  }

  private useStore(...args: InferStoreArgs<O>): InferStoreFromOptions<O> {
    const path = this.options.buildPath(...args);

    return fetchMapValue(this.stores, path, () => {
      const define = defineStore(path, () => this.options.setup(...args));
      return define() as InferStoreFromOptions<O>;
    });
  }

  private disposeStores() {
    const pinia = getActivePinia()!;

    for (const store of this.stores.values()) {
      destroyStore(store as Store, pinia);
    }

    this.stores.clear();
  }
}
