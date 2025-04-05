import { reactive, ref } from 'vue';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type AsyncFn = (...args: any[]) => Promise<void>;

export interface IAsyncAction<F extends AsyncFn> {
  call: F;
  isActive: boolean;
}

export function useAsyncAction<F extends AsyncFn>(fn: F): IAsyncAction<F> {
  const isActive = ref(false);

  const call = ((...args) => {
    isActive.value = true;
    return fn(...args).finally(() => isActive.value = false);
  }) as F;

  return reactive({
    call,
    isActive,
  }) as IAsyncAction<F>;
}
