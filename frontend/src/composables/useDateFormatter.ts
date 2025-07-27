import { computed, type ComputedRef, type MaybeRefOrGetter, toValue } from 'vue';

let formatter: Intl.DateTimeFormat;

export function useDateFormatter(inputRef: MaybeRefOrGetter<string | Date>): ComputedRef<string> {
  formatter ??= new Intl.DateTimeFormat(undefined, {
    dateStyle: 'medium',
    timeStyle: 'short',
  });

  const input = computed(() => {
    const input = toValue(inputRef);
    return typeof input === 'string' ? new Date(input) : input;
  });

  return computed(() => formatter.format(input.value));
}
