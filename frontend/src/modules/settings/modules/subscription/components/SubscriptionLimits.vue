<template>
  <Card title="Обмеження Підписки" class="subscription-limits">
    <div class="subscription-limits__list">
      <SubscriptionLimit
        v-for="limit in subscriptionStore.limits"
        :key="limit.key"
        :limit
      />
    </div>
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
    padding: 16px;
  }

  .subscription-limits__list {
    margin-top: 16px;
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
}
</style>
