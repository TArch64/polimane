import { computed, reactive } from 'vue';
import { useSessionStore } from '@/stores';
import { SubscriptionLimit, type UserLimit } from '@/enums';
import type { ICounter } from './counters';

export interface IUserCounter extends ICounter {
  max: number;
}

function useUserCounter(name: UserLimit): IUserCounter {
  const sessionStore = useSessionStore();

  const current = computed(() => sessionStore.subscription.counters[name]);
  const max = computed(() => sessionStore.plan.limits[name]!);
  const isReached = computed(() => current.value >= max.value);
  const overflowed = computed(() => Math.max(0, current.value - max.value));
  const isOverflowed = computed(() => overflowed.value > 0);

  function willOverflow(value: number): boolean {
    return current.value + value > max.value;
  }

  return reactive({
    isReached,
    isOverflowed,
    willOverflow,
    overflowed,
    current,
    max,
  });
}

export function useSchemasCreatedCounter() {
  return useUserCounter(SubscriptionLimit.SCHEMAS_CREATED);
}
