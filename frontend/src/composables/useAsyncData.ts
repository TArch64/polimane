import { nextTick, reactive, type Ref, ref, type UnwrapRef } from 'vue';
import { useAsyncState } from '@vueuse/core';
import { cloneReactiveAsRaw } from '@/helpers';
import { type IRouteTransition, useRouteTransition } from './useRouteTransition';

export interface IAsyncDataOptions<V> {
  loader: (current: V) => Promise<V>;
  default: V;
  once?: boolean;
}

export type OptimisticModify<D> = (data: D) => void;
export type OptimisticExecute = () => Promise<void>;

interface IOptimisticUpdateOptions<D> {
  state: Ref<D>;
  routeTransition: IRouteTransition;
}

export class OptimisticUpdate<D> {
  private state;
  private routeTransition;
  private isTransition = false;
  private modify!: OptimisticModify<D>;
  private temp: D | null = null;

  constructor(options: IOptimisticUpdateOptions<D>) {
    this.state = options.state;
    this.routeTransition = options.routeTransition;
  }

  inTransition(): OptimisticUpdate<D> {
    this.isTransition = true;
    return this;
  }

  begin(modify: OptimisticModify<D>): OptimisticUpdate<D> {
    this.modify = modify;
    return this;
  }

  async commit(execute: OptimisticExecute): Promise<void> {
    await this.withTransition(() => {
      this.temp = this.state.value;

      const value = cloneReactiveAsRaw(this.state.value);
      this.modify(value);
      this.state.value = value;
    });

    try {
      await execute();
      this.temp = null;
    } catch (error) {
      await this.withTransition(() => {
        this.state.value = this.temp!;
        this.temp = null;
      });

      throw error;
    }
  }

  private async withTransition(fn: () => void): Promise<void> {
    if (!this.isTransition) {
      return fn();
    }

    return this.routeTransition.start(() => {
      fn();
      return nextTick();
    });
  }
}

export interface IAsyncData<D> {
  data: D;
  load: () => Promise<void>;
  reset: () => void;
  optimisticUpdate: () => OptimisticUpdate<D>;
  isInitial: boolean;
  isLoading: boolean;
}

export function useAsyncData<D>(options: IAsyncDataOptions<D>): IAsyncData<D> {
  const routeTransition = useRouteTransition();
  const isInitial = ref(true);

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

  function optimisticUpdate(): OptimisticUpdate<D> {
    return new OptimisticUpdate<D>({
      state: state as Ref<D>,
      routeTransition,
    });
  }

  return reactive({
    data: state,
    isInitial,
    isLoading,
    reset,
    load,
    optimisticUpdate,
  }) as IAsyncData<D>;
}
