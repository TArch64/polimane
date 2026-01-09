import { defineStore } from 'pinia';
import { computed } from 'vue';
import { useSessionStore } from '@/stores';
import type { ISubscriptionCounters, ISubscriptionLimits, IUserSubscription } from '@/models';
import { getObjectEntries } from '@/helpers';

type LimitKey = keyof ISubscriptionLimits;

export enum SubscriptionLimitType {
  USER = 'user',
  FEATURE = 'feature',
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
    type: SubscriptionLimitType.USER,
    title: 'Створено Схем',
  },

  schemaBeads: {
    type: SubscriptionLimitType.FEATURE,
    title: 'Кількість Бісеру в Схемі',
  },

  sharedAccess: {
    type: SubscriptionLimitType.FEATURE,
    title: 'Користувачі з доступом до Схеми',
  },
};

type LimitFactory = (subscription: IUserSubscription, key: LimitKey, config: ILimitConfig) => ISubscriptionLimit;

const createUserLimit: LimitFactory = (subscription, key, config) => ({
  key,
  max: subscription.limits[key] ?? null,
  used: subscription.counters[key as keyof ISubscriptionCounters] ?? null,
  ...config,
});

const createFeatureLimit: LimitFactory = (subscription, key, config) => ({
  key,
  max: subscription.limits[key] ?? null,
  used: null,
  ...config,
});

const limitFactories: Record<SubscriptionLimitType, LimitFactory> = {
  [SubscriptionLimitType.USER]: createUserLimit,
  [SubscriptionLimitType.FEATURE]: createFeatureLimit,
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
