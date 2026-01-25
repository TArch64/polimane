<template>
  <Card class="subscription-plans">
    <SubscriptionPlan
      v-for="plan of plansStore.plans"
      :key="plan.id"
      :plan="plan"
      class="subscription-plans__item"
    />
  </Card>
</template>

<script setup lang="ts">
import { Card } from '@/components/card';
import { usePlansStore } from '@/stores';
import SubscriptionPlan from './SubscriptionPlan.vue';

const plansStore = usePlansStore();
</script>

<style scoped>
@layer page {
  .subscription-plans {
    padding: 0;
    display: flex;
    min-width: 0;
  }

  .subscription-plans__item {
    flex-basis: 0;
    flex-grow: 1;
    min-width: 0;
  }

  @media (max-width: 768px) {
    .subscription-plans {
      flex-direction: column-reverse;
    }
  }
}
</style>
