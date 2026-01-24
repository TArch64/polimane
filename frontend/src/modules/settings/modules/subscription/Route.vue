<template>
  <SubscriptionInfo />
  <SubscriptionLimits />
  <SubscriptionPlans v-if="plansStore.plans.length" />
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { SubscriptionInfo, SubscriptionLimits, SubscriptionPlans } from './components';
import { usePlansStore } from './stores';

defineOptions({
  beforeRouteEnter: definePreload<'settings-subscription'>(async () => {
    await usePlansStore().load();
  }),
});

const plansStore = usePlansStore();
</script>
