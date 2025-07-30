import { type MaybeRefOrGetter, nextTick, reactive, ref, toValue, watch } from 'vue';

export interface IFormData<D extends object> {
  data: D;
  hasChanges: boolean;
  reset: () => void;
}

export function useFormData<D extends object>(initial: MaybeRefOrGetter<D>): IFormData<D> {
  const hasChanges = ref(false);
  const data = ref<D>({ ...toValue(initial) });

  watch(data, () => {
    hasChanges.value = true;
  }, { deep: true });

  function reset(): void {
    data.value = { ...toValue(initial) };
    nextTick(() => hasChanges.value = false);
  }

  return reactive({
    data,
    hasChanges,
    reset,
  }) as IFormData<D>;
}
