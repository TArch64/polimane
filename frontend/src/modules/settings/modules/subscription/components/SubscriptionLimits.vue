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
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 16px;
  }
}
</style>
