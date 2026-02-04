<template>
  <p class="plan-limit">
    <span>{{ config.title }}</span>

    <span class="plan-limit__value">
      {{ formattedLimit }}
    </span>
  </p>
</template>

<script setup lang="ts">
import type { ISubscriptionPlan } from '@/models';
import { SubscriptionLimit } from '@/enums';
import type { IPlanLimitConfig } from '@/config';
import { useLimitFormatter } from '@/composables/subscription';

const props = defineProps<{
  plan: ISubscriptionPlan;
  limit: SubscriptionLimit;
  config: IPlanLimitConfig;
}>();

const formattedLimit = useLimitFormatter({
  plan: () => props.plan,
  limit: () => props.limit,
});
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
