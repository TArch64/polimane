import { defineStore } from 'pinia';
import { computed } from 'vue';
import { useSessionStore } from '@/stores';
import type { IUserSubscription, SubscriptionLimits, UserCounters } from '@/models';
import { getObjectEntries } from '@/helpers';
import { SubscriptionLimitType, SubscriptionPlanId } from '@/enums';
import { type HttpBody, useHttpClient } from '@/composables';
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
  max: subscription.plan.limits[key] ?? null,
  used: subscription.counters[key as keyof UserCounters] ?? null,
  ...config,
});

const createFeatureLimit: LimitFactory = (subscription, key, config) => ({
  key,
  max: subscription.plan.limits[key] ?? null,
  used: null,
  ...config,
});

const limitFactories: Record<SubscriptionLimitType, LimitFactory> = {
  [SubscriptionLimitType.USER]: createUserLimit,
  [SubscriptionLimitType.FEATURE]: createFeatureLimit,
};

interface IChangePlanBody {
  planId: SubscriptionPlanId;
}

export const useSubscriptionStore = defineStore('settings/subscription', () => {
  const http = useHttpClient();
  const sessionStore = useSessionStore();

  const limits = computed(() => {
    const limits: ISubscriptionLimit[] = [];

    for (const [key, config] of getObjectEntries(PLAN_LIMIT_CONFIGS)) {
      const createLimit = limitFactories[config.type];
      limits.push(createLimit(sessionStore.subscription, key, config));
    }

    return limits;
  });

  async function changePlan(planId: SubscriptionPlanId): Promise<void> {
    await http.post<HttpBody, IChangePlanBody>('/users/current/subscription/change', {
      planId,
    });

    try {
      await sessionStore.refresh();
    } catch {
      location.reload();
    }
  }

  return { limits, changePlan };
});
