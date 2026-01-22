import { computed, type MaybeRefOrGetter, reactive, toValue } from 'vue';
import { useSessionStore } from '@/stores';
import { type SchemaLimit, SubscriptionLimit } from '@/enums';
import type { ISchema } from '@/models';

export interface ISchemaLimit {
  isReached: boolean;
  current: number;
  max?: number;
}

function useSchemaLimit(name: SchemaLimit, schemaRef: MaybeRefOrGetter<ISchema>): ISchemaLimit {
  const sessionStore = useSessionStore();
  const subscription = computed(() => sessionStore.user.subscription);
  const schema = computed(() => toValue(schemaRef));

  const current = computed({
    get: () => schema.value.counters[name],
    set: (value: number) => schema.value.counters[name] = value,
  });

  const max = computed(() => subscription.value.limits[name]);
  const isReached = computed(() => max.value !== undefined && current.value >= max.value);

  return reactive({
    isReached,
    current,
    max,
  });
}

export function useSchemaBeadsLimit(schemaRef: MaybeRefOrGetter<ISchema>) {
  return useSchemaLimit(SubscriptionLimit.SCHEMA_BEADS, schemaRef);
}

export function useSharedAccessLimit(schemaRef: MaybeRefOrGetter<ISchema>) {
  return useSchemaLimit(SubscriptionLimit.SHARED_ACCESS, schemaRef);
}
