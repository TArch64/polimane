import { reactive, ref, type UnwrapRef } from 'vue';
import { useAsyncState } from '@vueuse/core';

export interface IAsyncDataOptions<V> {
  loader: () => Promise<V>;
  default: V;
  once?: boolean;
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

  const { state, isLoading, execute } = useAsyncState(async (): Promise<D> => {
    if (options.once && !isInitial.value) {
      return state.value as D;
    }

    try {
      return await options.loader();
    } finally {
      isInitial.value = false;
    }
  }, options.default, {
    immediate: options.immediate ?? false,
    throwError: true,
    shallow: false,
  });

  async function load() {
    await execute();
  }

  function setOptimisticUpdate(optimisticData: D): void {
    temp = state.value;
    state.value = optimisticData as UnwrapRef<D>;
  }

  function makeOptimisticUpdate(transform: OptimisticModify<D>): void {
    setOptimisticUpdate(transform(state.value as D));
  }

  function commitOptimisticUpdate(): void {
    temp = null;
  }

  function rollbackOptimisticUpdate(): void {
    state.value = temp!;
    temp = null;
  }

  return reactive({
    data: state,
    isInitial,
    isLoading,
    load,
    setOptimisticUpdate,
    makeOptimisticUpdate,
    commitOptimisticUpdate,
    rollbackOptimisticUpdate,
  }) as IAsyncData<D>;
}
