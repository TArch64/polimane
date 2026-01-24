import { defineStore } from 'pinia';
import { computed } from 'vue';
import { useSessionStore } from '@/stores';
import type { IUserSubscription, SubscriptionLimits, UserCounters } from '@/models';
import { getObjectEntries } from '@/helpers';
import { SubscriptionLimitType } from '@/enums';
import { type IPlanLimitConfig, PLAN_LIMIT_CONFIGS } from './plansStore';

type LimitKey = keyof SubscriptionLimits;

export interface ISubscriptionLimit {
  key: LimitKey;
  type: SubscriptionLimitType;
  title: string;
  max: number | null;
  used: number | null;
}

type LimitFactory = (subscription: IUserSubscription, key: LimitKey, config: IPlanLimitConfig) => ISubscriptionLimit;

const createUserLimit: LimitFactory = (subscription, key, config) => ({
  key,
  max: subscription.limits[key] ?? null,
  used: subscription.counters[key as keyof UserCounters] ?? null,
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

    for (const [key, config] of getObjectEntries(PLAN_LIMIT_CONFIGS)) {
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
