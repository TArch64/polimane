<template>
  <LimitView>
    <LimitInfo
      :title="limit.title"
      :value="`${limit.used}/${limit.max}`"
    />

    <LimitProgress :value="usage" />
  </LimitView>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { ISubscriptionLimit } from '../../stores';
import { LimitInfo, LimitProgress, LimitView } from './base';

const props = defineProps<{
  limit: ISubscriptionLimit;
}>();

const usage = computed(() => {
  if (props.limit.max === null) {
    return 0;
  }

  const used = Math.min(props.limit.used!, props.limit.max);
  return (used / props.limit.max) * 100;
});
</script>
