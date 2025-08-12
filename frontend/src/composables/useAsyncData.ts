import { reactive, ref, type UnwrapRef } from 'vue';
import { useAsyncState } from '@vueuse/core';

export interface IAsyncDataOptions<V> {
  loader: () => Promise<V>;
  default: V;
  immediate?: boolean;
}

export type OptimisticModify<D> = (data: D) => D;

export interface IAsyncData<D> {
  data: D;
  load: () => Promise<void>;
  makeOptimisticUpdate: (transform: OptimisticModify<D>) => void;
  setOptimisticUpdate: (data: D) => void;
  commitOptimisticUpdate: () => void;
  rollbackOptimisticUpdate: () => void;
  isInitial: boolean;
  isLoading: boolean;
}

export function useAsyncData<D>(options: IAsyncDataOptions<D>): IAsyncData<D> {
  const isInitial = ref(true);
  let temp: UnwrapRef<D> | null = null;

  const { state: data, isLoading, execute } = useAsyncState(() => {
    return options
      .loader()
      .finally(() => isInitial.value = false);
  }, options.default, {
    immediate: options.immediate ?? false,
    throwError: true,
    shallow: false,
  });

  async function load() {
    await execute();
  }

  function setOptimisticUpdate(optimisticData: D): void {
    temp = data.value;
    data.value = optimisticData as UnwrapRef<D>;
  }

  function makeOptimisticUpdate(transform: OptimisticModify<D>): void {
    setOptimisticUpdate(transform(data.value as D));
  }

  function commitOptimisticUpdate(): void {
    temp = null;
  }

  function rollbackOptimisticUpdate(): void {
    data.value = temp!;
    temp = null;
  }

  return reactive({
    data,
    isInitial,
    isLoading,
    load,
    setOptimisticUpdate,
    makeOptimisticUpdate,
    commitOptimisticUpdate,
    rollbackOptimisticUpdate,
  }) as IAsyncData<D>;
}
