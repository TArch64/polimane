<template>
  <p class="upgrading-plan-limit">
    <span>{{ config.title }}</span>

    <span class="upgrading-plan-limit__value">
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
  .upgrading-plan-limit {
    display: flex;
    justify-content: space-between;
  }

  .upgrading-plan-limit__value {
    font-weight: 500;
  }
}
</style>
