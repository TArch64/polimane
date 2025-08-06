import { computed, type ComputedRef, type MaybeRefOrGetter, toValue } from 'vue';
import type { ISchemaRow, SchemaPattern } from '@/models';
import { useRowsStore } from '../stores';
import { useObjectParent } from '../models';

export function useRowTitle(rowRef: MaybeRefOrGetter<ISchemaRow>): ComputedRef<string> {
  const pattern = useObjectParent<SchemaPattern>(rowRef);
  const rowsStore = useRowsStore(pattern);

  const index = computed(() => rowsStore.rows.indexOf(toValue(rowRef)));
  return computed(() => `Рядок #${index.value + 1}`);
}
