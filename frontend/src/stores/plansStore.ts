import { defineStore } from 'pinia';
import { toRef } from 'vue';
import { useAsyncData, useHttpClient } from '@/composables';
import type { ISubscriptionPlan } from '@/models';

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
