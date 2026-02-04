import { computed, type MaybeRefOrGetter, reactive, toValue } from 'vue';
import { useSessionStore } from '@/stores';
import { type SchemaLimit, SubscriptionLimit } from '@/enums';
import type { ISchemaCountersLike } from '@/models';
import type { ICounter } from './counters';

export interface ISchemaCounter extends ICounter {
}

function useSchemaCounter(name: SchemaLimit, schemaRef: MaybeRefOrGetter<ISchemaCountersLike>): ISchemaCounter {
  const sessionStore = useSessionStore();
  const counters = computed(() => toValue(schemaRef).counters);

  const current = computed({
    get: () => counters.value[name],
    set: (value: number) => counters.value[name] = value,
  });

  const max = computed(() => sessionStore.getLimit(name));

  const overflowed = computed(() => {
    if (max.value === undefined) return 0;
    return Math.max(0, current.value - max.value);
  });

  const isReached = computed(() => max.value !== undefined && current.value >= max.value);
  const isOverflowed = computed(() => overflowed.value > 0);

  function willOverflow(value: number): boolean {
    return max.value !== undefined && current.value + value > max.value;
  }

  return reactive({
    isReached,
    isOverflowed,
    current,
    max,
    willOverflow,
    overflowed,
  });
}

export function useSchemaBeadsCounter(schemaRef: MaybeRefOrGetter<ISchemaCountersLike>) {
  return useSchemaCounter(SubscriptionLimit.SCHEMA_BEADS, schemaRef);
}

export function useSchemasSharedAccessCounter(schemaRef: MaybeRefOrGetter<ISchemaCountersLike>) {
  return useSchemaCounter(SubscriptionLimit.SHARED_ACCESS, schemaRef);
}
