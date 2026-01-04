import { defineStore } from 'pinia';
import { computed } from 'vue';
import { useSessionStore } from '@/stores';
import type { ISubscriptionCounters, ISubscriptionLimits, IUserSubscription } from '@/models';
import { getObjectKeys } from '@/helpers';

type LimitKey = keyof ISubscriptionLimits;

export enum SubscriptionLimitType {
  COUNTER = 'counter',
  PER_FEATURE = 'per_feature',
}

export interface ISubscriptionLimit {
  key: LimitKey;
  type: SubscriptionLimitType;
  max: number | null;
  used: number | null;
}

const LIMIT_KEYS = {
  schemasCreated: true,
  sharedAccess: true,
} as const satisfies Record<LimitKey, boolean>;

const COUNTER_LIMIT_KEYS: LimitKey[] = [
  'schemasCreated',
];

function getLimitType(limitKey: LimitKey): SubscriptionLimitType {
  return COUNTER_LIMIT_KEYS.includes(limitKey)
    ? SubscriptionLimitType.COUNTER
    : SubscriptionLimitType.PER_FEATURE;
}

type LimitFactory = (subscription: IUserSubscription, key: LimitKey) => ISubscriptionLimit;

const createCounterLimit: LimitFactory = (subscription, key) => ({
  key,
  type: SubscriptionLimitType.COUNTER,
  max: subscription.limits[key] ?? null,
  used: subscription.counters[key as keyof ISubscriptionCounters] ?? null,
});

const createPerFeatureLimit: LimitFactory = (subscription, key) => ({
  key,
  type: SubscriptionLimitType.PER_FEATURE,
  max: subscription.limits[key] ?? null,
  used: null,
});

const limitFactories: Record<SubscriptionLimitType, LimitFactory> = {
  [SubscriptionLimitType.COUNTER]: createCounterLimit,
  [SubscriptionLimitType.PER_FEATURE]: createPerFeatureLimit,
};

export const useSubscriptionStore = defineStore('settings/subscription', () => {
  const sessionStore = useSessionStore();
  const subscription = computed(() => sessionStore.user.subscription);

  const limits = computed(() => {
    const limits: ISubscriptionLimit[] = [];

    for (const limitKey of getObjectKeys(LIMIT_KEYS)) {
      const type = getLimitType(limitKey);
      const createLimit = limitFactories[type];
      limits.push(createLimit(subscription.value, limitKey));
    }

    return limits;
  });

  return {
    subscription,
    limits,
  };
});
