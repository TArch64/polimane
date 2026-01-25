<template>
  <SubscriptionInfo />
  <SubscriptionLimits />
  <SubscriptionPlans v-if="plansStore.plans.length" />
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { usePlansStore } from '@/stores';
import { SubscriptionInfo, SubscriptionLimits, SubscriptionPlans } from './components';

defineOptions({
  beforeRouteEnter: definePreload<'settings-subscription'>(async () => {
    await usePlansStore().load();
  }),
});

const plansStore = usePlansStore();
</script>
