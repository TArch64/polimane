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

const used = computed(() => props.limit.used!);
const max = computed(() => props.limit.max);

const usage = computed(() => {
  if (max.value === null) return 0;
  return Math.min((used.value / max.value) * 100, 100);
});
</script>
