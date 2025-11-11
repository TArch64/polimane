import { nextTick, reactive, ref, type UnwrapRef } from 'vue';
import { useAsyncState } from '@vueuse/core';
import { useRouteTransition } from './useRouteTransition';

export interface IAsyncDataOptions<V> {
  loader: (current: V) => Promise<V>;
  default: V;
  once?: boolean;
}

export type OptimisticModify<D> = (data: D) => D;
export type OptimisticExecute = () => Promise<void>;

export interface IOptimisticOptions {
  transition?: boolean;
}

export interface IAsyncData<D> {
  data: D;
  load: () => Promise<void>;
  reset: () => void;
  makeOptimisticUpdate: (transform: OptimisticModify<D>, options?: IOptimisticOptions) => void;
  executeOptimisticUpdate: (execute: OptimisticExecute, options?: IOptimisticOptions) => Promise<void>;
  isInitial: boolean;
  isLoading: boolean;
}

export function useAsyncData<D>(options: IAsyncDataOptions<D>): IAsyncData<D> {
  const routeTransition = useRouteTransition();

  const isInitial = ref(true);
  let temp: UnwrapRef<D> | null = null;

  const { state, isLoading, execute } = useAsyncState(async (): Promise<D> => {
    try {
      return await options.loader(state.value as D);
    } finally {
      isInitial.value = false;
    }
  }, options.default, {
    immediate: false,
    resetOnExecute: false,
    throwError: true,
    shallow: false,
  });

  function reset() {
    state.value = options.default as UnwrapRef<D>;
    isInitial.value = true;
  }

  async function load() {
    if (options.once && !isInitial.value) {
      return;
    }

    await execute();
  }

  function withTransition(enabled: boolean | undefined, fn: () => void): void {
    if (!enabled) {
      return fn();
    }

    routeTransition.start(() => {
      fn();
      return nextTick();
    });
  }

  function makeOptimisticUpdate(transform: OptimisticModify<D>, options: IOptimisticOptions = {}): void {
    withTransition(options.transition, () => {
      temp = state.value;
      state.value = transform(state.value as D) as UnwrapRef<D>;
    });
  }

  async function executeOptimisticUpdate(execute: OptimisticExecute, options: IOptimisticOptions = {}): Promise<void> {
    try {
      await execute();
      temp = null;
    } catch (error) {
      withTransition(options.transition, () => {
        state.value = temp!;
        temp = null;
      });

      throw error;
    }
  }

  return reactive({
    data: state,
    isInitial,
    isLoading,
    reset,
    load,
    makeOptimisticUpdate,
    executeOptimisticUpdate,
  }) as IAsyncData<D>;
}
