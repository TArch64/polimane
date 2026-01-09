import { computed, reactive } from 'vue';
import { useSessionStore } from '@/stores';
import { SubscriptionLimit, type UserLimit } from '@/enums';

export interface IUserLimit {
  isReached: boolean;
}

function useUserLimit(name: UserLimit): IUserLimit {
  const sessionStore = useSessionStore();
  const subscription = computed(() => sessionStore.user.subscription);

  const counter = computed(() => subscription.value.counters[name]);
  const limit = computed(() => subscription.value.limits[name]!);
  const isReached = computed(() => counter.value >= limit.value);

  return reactive({ isReached });
}

export function useSchemasCreatedLimit() {
  return useUserLimit(SubscriptionLimit.SCHEMAS_CREATED);
}
