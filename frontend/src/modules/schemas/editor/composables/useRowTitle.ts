import { computed, type ComputedRef, type MaybeRefOrGetter, toValue } from 'vue';
import type { ISchemaPattern, ISchemaRow } from '@/models';
import { useRowsStore } from '../stores';
import { getObjectParent } from '../models';

export function useRowTitle(rowRef: MaybeRefOrGetter<ISchemaRow>): ComputedRef<string> {
  const row = computed(() => toValue(rowRef));
  const pattern = computed(() => getObjectParent<ISchemaPattern>(row.value));
  const rowsStore = useRowsStore(pattern);

  const index = computed(() => rowsStore.rows.indexOf(row.value));
  return computed(() => `Рядок #${index.value + 1}`);
}
