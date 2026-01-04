<template>
  <Card class="subscription-limits">
    <SubscriptionLimit
      v-for="(_, limitKey) in limitKeys"
      :key="limitKey"
      :limitKey
    />
  </Card>
</template>

<script setup lang="ts">
import type { ISubscriptionLimits } from '@/models';
import { Card } from '@/components/card';
import SubscriptionLimit from './SubscriptionLimit.vue';

const limitKeys = {
  schemasCreated: true,
  sharedAccess: true,
} as const satisfies Record<keyof ISubscriptionLimits, boolean>;
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
