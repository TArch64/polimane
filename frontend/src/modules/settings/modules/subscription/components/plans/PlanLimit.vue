<template>
  <p class="plan-limit">
    <span>{{ config.title }}</span>

    <span class="plan-limit__value">
      {{ formattedLimit || INFINITY_SYMBOL }}
    </span>
  </p>
</template>

<script setup lang="ts">
import { useNumberFormatter } from '@/composables';
import type { ISubscriptionPlan } from '@/models';
import { SubscriptionLimit } from '@/enums';
import { INFINITY_SYMBOL } from '@/config';
import type { IPlanLimitConfig } from '../../stores';

const props = defineProps<{
  plan: ISubscriptionPlan;
  limit: SubscriptionLimit;
  config: IPlanLimitConfig;
}>();

const formattedLimit = useNumberFormatter(() => props.plan.limits[props.limit]);
</script>

<style scoped>
@layer page {
  .plan-limit {
    display: flex;
    justify-content: space-between;
    font-size: var(--font-sm);
  }

  .plan-limit__value {
    font-weight: 500;
  }
}
</style>
