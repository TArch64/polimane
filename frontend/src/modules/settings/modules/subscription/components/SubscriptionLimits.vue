<template>
  <Card class="subscription-limits">
    <SubscriptionLimit
      v-for="limit in subscriptionStore.limits"
      :key="limit.key"
      :limit
    />
  </Card>
</template>

<script setup lang="ts">
import { Card } from '@/components/card';
import { useSubscriptionStore } from '../stores';
import { SubscriptionLimit } from './limit';

const subscriptionStore = useSubscriptionStore();
</script>

<style scoped>
@layer page {
  .subscription-limits {
    display: grid;
    grid-template-columns: repeat(auto-fit, 140px);
    gap: 32px 24px;
    padding: 28px 16px 36px;
    place-items: center;
    justify-content: space-evenly;
  }

  .subscription-limits__holder {
    display: none;
  }

  @media (min-width: 900px) {
    .subscription-limits__holder {
      display: block;
    }
  }
}
</style>
