<template>
  <div>
    <p>{{ limitKey }}</p>
    <p v-if="counter === null">{{ limit }}</p>
    <p v-else>{{ counter }} / {{ limit }}</p>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { ISubscriptionCounters, ISubscriptionLimits } from '@/models';
import { useSubscriptionStore } from '../stores';

const props = defineProps<{
  limitKey: keyof ISubscriptionLimits;
}>();

const subscriptionStore = useSubscriptionStore();

const limit = computed(() => {
  return subscriptionStore.subscription.limits[props.limitKey] ?? '-';
});

const counter = computed(() => {
  return subscriptionStore.subscription.counters[props.limitKey as keyof ISubscriptionCounters] ?? null;
});
</script>
