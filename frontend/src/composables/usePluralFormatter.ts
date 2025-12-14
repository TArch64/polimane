import { computed, type ComputedRef, type MaybeRefOrGetter, toValue } from 'vue';
import { LOCALE } from '@/config';
import type { MarkOptional } from '@/types';

export type PluralMapping = MarkOptional<Record<Intl.LDMLPluralRule, string>, 'zero' | 'two' | 'many'>;

export function usePluralFormatter(inputRef: MaybeRefOrGetter<number>, mapping: PluralMapping): ComputedRef<string> {
  const formatter = new Intl.PluralRules(LOCALE);
  const input = computed(() => toValue(inputRef));
  const category = computed(() => formatter.select(input.value));
  return computed(() => mapping[category.value] ?? mapping.other);
}

export const SCHEMA_PLURAL: PluralMapping = {
  one: 'схема',
  few: 'схеми',
  other: 'схем',
};
