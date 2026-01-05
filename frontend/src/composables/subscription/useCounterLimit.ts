import { computed, reactive } from 'vue';
import { useSessionStore } from '@/stores';
import type { ISubscriptionCounters, ISubscriptionLimits } from '@/models';

type CounterName = keyof ISubscriptionCounters & keyof ISubscriptionLimits;

export interface ICounterLimit {
  value: number;
  limit: number;
  isReached: boolean;
  increase: (delta?: number) => void;
  decrease: (delta?: number) => void;
}

function useCounterLimit(name: CounterName): ICounterLimit {
  const sessionStore = useSessionStore();
  const subscription = computed(() => sessionStore.user.subscription);
  const value = computed(() => subscription.value.counters[name]);
  const limit = computed(() => subscription.value.limits[name]!);
  const isReached = computed(() => value.value >= limit.value);

  function increase(delta = 1) {
    subscription.value.counters[name] += delta;
  }

  function decrease(delta = 1) {
    subscription.value.counters[name] -= delta;

    if (subscription.value.counters[name] < 0) {
      subscription.value.counters[name] = 0;
    }
  }

  return reactive({
    value,
    limit,
    isReached,
    increase,
    decrease,
  });
}

export function useSchemasCreatedCounterLimit() {
  return useCounterLimit('schemasCreated');
}
