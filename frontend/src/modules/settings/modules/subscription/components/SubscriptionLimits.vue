<template>
  <Card title="Обмеження Підписки" class="subscription-limits">
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
    display: flex;
    flex-direction: column;
    gap: 24px;
    padding: 16px;
  }
}
</style>
