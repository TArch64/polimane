import { reactive, ref } from 'vue';
import { useAsyncState } from '@vueuse/core';

export interface IAsyncDataOptions<V> {
  loader: () => Promise<V>;
  default: V;
  immediate?: boolean;
}

export interface IAsyncData<V> {
  data: V;
  load: () => Promise<void>;
  isInitial: boolean;
  isLoading: boolean;
}

export function useAsyncData<D>(options: IAsyncDataOptions<D>): IAsyncData<D> {
  const isInitial = ref(true);

  const { state: data, isLoading, execute } = useAsyncState(() => {
    return options
      .loader()
      .finally(() => isInitial.value = false);
  }, options.default, {
    immediate: options.immediate ?? false,
    throwError: true,
  });

  async function load() {
    await execute();
  }

  return reactive({ data, isInitial, isLoading, load });
}
