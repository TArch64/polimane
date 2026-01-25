import { computed, reactive } from 'vue';
import { useSessionStore } from '@/stores';
import { SubscriptionLimit, type UserLimit } from '@/enums';

export interface IUserCounter {
  isReached: boolean;
  willReach: (value: number) => boolean;
  current: number;
  max: number;
}

function useUserCounter(name: UserLimit): IUserCounter {
  const sessionStore = useSessionStore();

  const current = computed(() => sessionStore.subscription.counters[name]);
  const max = computed(() => sessionStore.plan.limits[name]!);
  const isReached = computed(() => current.value >= max.value);

  function willReach(value: number): boolean {
    return current.value + value > max.value;
  }

  return reactive({
    isReached,
    willReach,
    current,
    max,
  });
}

export function useSchemasCreatedCounter() {
  return useUserCounter(SubscriptionLimit.SCHEMAS_CREATED);
}
