<template>
  <LimitView
    :title="limit.title"
    :subtitle="`${limit.used}/${limit.max}`"
    :progress="usedPercentage"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { ISubscriptionLimit } from '../../stores';
import LimitView from './LimitView.vue';

const props = defineProps<{
  limit: ISubscriptionLimit;
}>();

const usedPercentage = computed(() => {
  if (props.limit.max === null) {
    return 0;
  }

  const used = Math.min(props.limit.used!, props.limit.max);
  return (used / props.limit.max) * 100;
});
</script>
