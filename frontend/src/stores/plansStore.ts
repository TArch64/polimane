import { defineStore } from 'pinia';
import { toRef } from 'vue';
import { useAsyncData, useHttpClient } from '@/composables';
import type { ISubscriptionPlan, SubscriptionLimits } from '@/models';
import { SubscriptionLimitType } from '@/enums';

type LimitKey = keyof SubscriptionLimits;

export interface IPlanLimitConfig {
  type: SubscriptionLimitType;
  title: string;
}

export const PLAN_LIMIT_CONFIGS: Record<LimitKey, IPlanLimitConfig> = {
  schemasCreated: {
    type: SubscriptionLimitType.USER,
    title: 'Схеми',
  },

  schemaBeads: {
    type: SubscriptionLimitType.FEATURE,
    title: 'Бісер в Схемі',
  },

  sharedAccess: {
    type: SubscriptionLimitType.FEATURE,
    title: 'Користувачі Схеми',
  },
};

export const usePlansStore = defineStore('plans', () => {
  const http = useHttpClient();

  const plans = useAsyncData({
    async loader() {
      const plans = await http.get<ISubscriptionPlan[]>('/users/current/subscription/plans');
      return plans.sort((a, b) => a.tier - b.tier);
    },

    once: true,
    default: [],
  });

  return {
    load: plans.load,
    plans: toRef(plans, 'data'),
  };
});
