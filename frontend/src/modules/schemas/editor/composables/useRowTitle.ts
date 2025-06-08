import { computed, type ComputedRef, type MaybeRefOrGetter, toValue } from 'vue';
import type { ISchemaPattern, ISchemaRow } from '@/models';
import { useRowsStore } from '../stores';
import { useObjectParent } from '../models';

export function useRowTitle(rowRef: MaybeRefOrGetter<ISchemaRow>): ComputedRef<string> {
  const pattern = useObjectParent<ISchemaPattern>(rowRef);
  const rowsStore = useRowsStore(pattern);

  const index = computed(() => rowsStore.rows.indexOf(toValue(rowRef)));
  return computed(() => `Рядок #${index.value + 1}`);
}
