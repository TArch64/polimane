import { ref, unref } from 'vue';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type AsyncFn = (...args: any[]) => Promise<void>;

export type AsyncAction<F extends AsyncFn> = F & {
  isActive: boolean;
};

export function useAsyncAction<F extends AsyncFn>(fn: F): AsyncAction<F> {
  const isActive = ref(false);

  const call = ((...args) => {
    isActive.value = true;
    return fn(...args).finally(() => isActive.value = false);
  }) as F;

  const extension = {
    isActive,
  };

  return new Proxy(call, {
    get: (target: F, property: string): unknown => {
      if (property in extension) {
        return unref(extension[property as keyof typeof extension]);
      }
      return Reflect.get(target, property);
    },
  }) as AsyncAction<F>;
}
