import { computed, type ComputedRef, type MaybeRefOrGetter, toValue } from 'vue';
import { LOCALE } from '@/config';
import type { MaybeValue } from '@/types';

type Formatter = (inputRef: MaybeRefOrGetter<MaybeValue<number>>) => ComputedRef<string>;

function createFormatter(options: Intl.NumberFormatOptions): Formatter {
  let formatter: Intl.NumberFormat;

  return (inputRef) => {
    formatter ??= new Intl.NumberFormat(LOCALE, options);
    const input = computed(() => toValue(inputRef));
    return computed(() => input.value ? formatter.format(input.value) : '');
  };
}

export const useNumberFormatter = createFormatter({});

export const useIntPercentageFormatter = createFormatter({
  style: 'percent',
  maximumFractionDigits: 0,
});

export const useCurrencyFormatter = createFormatter({
  style: 'currency',
  currency: 'UAH',
  maximumFractionDigits: 2,
  minimumFractionDigits: 0,
});
