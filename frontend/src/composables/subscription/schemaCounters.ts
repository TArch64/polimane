import { computed, type MaybeRefOrGetter, reactive, toValue } from 'vue';
import { useSessionStore } from '@/stores';
import { type SchemaLimit, SubscriptionLimit } from '@/enums';
import type { ISchema } from '@/models';

export interface ISchemaCounter {
  isReached: boolean;
  current: number;
  max?: number;
}

function useSchemaCounter(name: SchemaLimit, schemaRef: MaybeRefOrGetter<ISchema>): ISchemaCounter {
  const sessionStore = useSessionStore();
  const schema = computed(() => toValue(schemaRef));

  const current = computed({
    get: () => schema.value.counters[name],
    set: (value: number) => schema.value.counters[name] = value,
  });

  const max = computed(() => sessionStore.getLimit(name));
  const isReached = computed(() => max.value !== undefined && current.value >= max.value);

  return reactive({
    isReached,
    current,
    max,
  });
}

export function useSchemaBeadsCounter(schemaRef: MaybeRefOrGetter<ISchema>) {
  return useSchemaCounter(SubscriptionLimit.SCHEMA_BEADS, schemaRef);
}

export function useSharedAccessCounter(schemaRef: MaybeRefOrGetter<ISchema>) {
  return useSchemaCounter(SubscriptionLimit.SHARED_ACCESS, schemaRef);
}
