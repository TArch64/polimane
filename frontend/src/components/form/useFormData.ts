import { reactive, ref, watch } from 'vue';

export interface IFormData<D extends object> {
  data: D;
  hasChanges: boolean;
  reset: () => void;
}

export function useFormData<D extends object>(initial: D): IFormData<D> {
  const hasChanges = ref(false);
  const data = ref<D>({ ...initial });

  watch(data, () => {
    hasChanges.value = true;
  }, { deep: true });

  function reset(): void {
    data.value = { ...initial };
    hasChanges.value = false;
  }

  return reactive({
    data,
    hasChanges,
    reset,
  }) as IFormData<D>;
}
