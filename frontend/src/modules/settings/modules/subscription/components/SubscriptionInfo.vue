<template>
  <Card class="subscription-info">
    <h3 class="subscription-info__plan-title">
      {{ planTitle }} Підписка
    </h3>
  </Card>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { Card } from '@/components/card';
import { SubscriptionPlan } from '@/enums';
import { useSubscriptionStore } from '../stores';

const subscriptionStore = useSubscriptionStore();

const planTitles: Record<SubscriptionPlan, string> = {
  [SubscriptionPlan.BETA]: 'Бета',
  [SubscriptionPlan.BASIC]: 'Базова',
  [SubscriptionPlan.PRO]: 'Продвинута',
};

const planTitle = computed(() => planTitles[subscriptionStore.subscription.plan]);
</script>

<style scoped>
@layer page {
  .subscription-info {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 32px;
  }

  .subscription-info__plan-title {
    font-size: 28px;
  }
}
</style>
