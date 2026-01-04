import { defineStore } from 'pinia';
import { computed } from 'vue';
import { useSessionStore } from '@/stores';
import type { ISubscriptionCounters, ISubscriptionLimits, IUserSubscription } from '@/models';
import { getObjectEntries } from '@/helpers';

type LimitKey = keyof ISubscriptionLimits;

export enum SubscriptionLimitType {
  COUNTER = 'counter',
  PER_FEATURE = 'per_feature',
}

export interface ISubscriptionLimit {
  key: LimitKey;
  type: SubscriptionLimitType;
  title: string;
  max: number | null;
  used: number | null;
}

interface ILimitConfig {
  type: SubscriptionLimitType;
  title: string;
}

const LIMIT_KEYS: Record<LimitKey, ILimitConfig> = {
  schemasCreated: {
    type: SubscriptionLimitType.COUNTER,
    title: 'Створені Схеми',
  },

  sharedAccess: {
    type: SubscriptionLimitType.PER_FEATURE,
    title: 'Спільний доступ',
  },
};

type LimitFactory = (subscription: IUserSubscription, key: LimitKey, config: ILimitConfig) => ISubscriptionLimit;

const createCounterLimit: LimitFactory = (subscription, key, config) => ({
  key,
  max: subscription.limits[key] ?? null,
  used: subscription.counters[key as keyof ISubscriptionCounters] ?? null,
  ...config,
});

const createPerFeatureLimit: LimitFactory = (subscription, key, config) => ({
  key,
  max: subscription.limits[key] ?? null,
  used: null,
  ...config,
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

    for (const [key, config] of getObjectEntries(LIMIT_KEYS)) {
      const createLimit = limitFactories[config.type];
      limits.push(createLimit(subscription.value, key, config));
    }

    return limits;
  });

  return {
    subscription,
    limits,
  };
});
