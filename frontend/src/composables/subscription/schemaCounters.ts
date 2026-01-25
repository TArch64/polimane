import { computed, type MaybeRefOrGetter, reactive, toValue } from 'vue';
import { useSessionStore } from '@/stores';
import { type SchemaLimit, SubscriptionLimit } from '@/enums';
import type { ISchema } from '@/models';

export interface ISchemaCounter {
  isReached: boolean;
  current: number;
  max?: number;
  willOverlow: (value: number) => boolean;
}

function useSchemaCounter(name: SchemaLimit, schemaRef: MaybeRefOrGetter<ISchema>): ISchemaCounter {
  const sessionStore = useSessionStore();
  const counters = computed(() => toValue(schemaRef).counters);

  const current = computed({
    get: () => counters.value[name],
    set: (value: number) => counters.value[name] = value,
  });

  const max = computed(() => sessionStore.getLimit(name));
  const isReached = computed(() => max.value !== undefined && current.value >= max.value);

  function willOverlow(value: number): boolean {
    return max.value !== undefined && current.value + value > max.value;
  }

  return reactive({
    isReached,
    current,
    max,
    willOverlow,
  });
}

export function useSchemaBeadsCounter(schemaRef: MaybeRefOrGetter<ISchema>) {
  return useSchemaCounter(SubscriptionLimit.SCHEMA_BEADS, schemaRef);
}

export function useSharedAccessCounter(schemaRef: MaybeRefOrGetter<ISchema>) {
  return useSchemaCounter(SubscriptionLimit.SHARED_ACCESS, schemaRef);
}
