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
import { getSubscriptionPlanName } from '@/enums';
import { useSubscriptionStore } from '../stores';

const subscriptionStore = useSubscriptionStore();

const planTitle = computed(() => getSubscriptionPlanName(subscriptionStore.subscription.planId));
</script>

<style scoped>
@layer page {
  .subscription-info {
    display: flex;
    align-items: center;
    justify-content: space-evenly;
    padding: 32px;
  }

  .subscription-info__plan-title {
    font-size: var(--font-xlg);
  }

  @media (max-width: 768px) {
    .subscription-info__plan-title {
      font-size: var(--font-lg);
    }
  }
}
</style>
