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

export const usePlansStore = defineStore('settings/subscription/plans', () => {
  const http = useHttpClient();

  const plans = useAsyncData({
    async loader() {
      return http.get<ISubscriptionPlan[]>('/users/current/plans');
    },

    once: true,
    default: [],
  });

  async function load(): Promise<void> {
    await plans.load();
  }

  return {
    load,
    plans: toRef(plans, 'data'),
  };
});
