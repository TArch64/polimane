import { computed, type ComputedRef, type MaybeRefOrGetter, toValue } from 'vue';
import { LOCALE } from '@/config';
import type { MaybeValue } from '@/types';

let formatter: Intl.NumberFormat;

export function useNumberFormatter(inputRef: MaybeRefOrGetter<MaybeValue<number>>): ComputedRef<string> {
  formatter ??= new Intl.NumberFormat(LOCALE);
  const input = computed(() => toValue(inputRef));
  return computed(() => input.value ? formatter.format(input.value) : '');
}
